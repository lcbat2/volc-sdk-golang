package producer

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	. "github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
)

type SDKProducerTestSuite struct {
	suite.Suite

	cli      Client
	producer Producer
	project  string
	topic    string
}

func newClientWithEnv() Client {
	return NewClient(os.Getenv("LOG_SERVICE_ENDPOINT"), os.Getenv("LOG_SERVICE_AK"),
		os.Getenv("LOG_SERVICE_SK"), "", os.Getenv("LOG_SERVICE_REGION"))
}

func createProject(projectName, description, region string, cli Client) (string, error) {
	createProjectReq := &CreateProjectRequest{
		ProjectName: projectName,
		Description: description,
		Region:      region,
	}

	createProjectResp, err := cli.CreateProject(createProjectReq)

	if err != nil {
		return "", err
	}

	return createProjectResp.ProjectID, nil
}

func createTopic(projectId, topicName, description string, shardCount int, ttl uint16, cli Client) (string, error) {
	createTopicReq := &CreateTopicRequest{
		ProjectID:   projectId,
		TopicName:   topicName,
		Ttl:         ttl,
		Description: description,
		ShardCount:  shardCount,
	}

	createTopicResp, err := cli.CreateTopic(createTopicReq)
	if err != nil {
		return "", err
	}

	return createTopicResp.TopicID, nil
}

func createIndex(topicID string, fulltextInfo *FullTextInfo, KeyValue *[]KeyValueInfo, cli Client) error {
	createIndexReq := &CreateIndexRequest{
		TopicID:  topicID,
		FullText: fulltextInfo,
		KeyValue: KeyValue,
	}
	_, err := cli.CreateIndex(createIndexReq)
	if err != nil {
		return err
	}

	time.Sleep(time.Minute)

	return nil
}

func (suite *SDKProducerTestSuite) SetupTest() {
	suite.cli = newClientWithEnv()

	projectId, err := createProject("golang-sdk-create-topic-"+uuid.New().String(), "test",
		os.Getenv("LOG_SERVICE_REGION"), suite.cli)
	suite.NoError(err)
	suite.project = projectId

	topicId, err := createTopic(projectId, "golang-sdk-create-index-"+uuid.New().String(),
		"test", 1, 1, suite.cli)
	suite.NoError(err)
	suite.topic = topicId

	keyValueList := make([]KeyValueInfo, 0)
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-1",
		Value: Value{
			ValueType:      "text",
			Delimiter:      "",
			CasSensitive:   false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	keyValueList = append(keyValueList, KeyValueInfo{
		Key: "key-2",
		Value: Value{
			ValueType:      "long",
			Delimiter:      "",
			CasSensitive:   false,
			IncludeChinese: false,
			SQLFlag:        true,
		},
	})
	suite.NoError(createIndex(topicId, nil, &keyValueList, suite.cli))

	// init producer
	producerCfg := GetDefaultProducerConfig()
	producerCfg.Endpoint = os.Getenv("LOG_SERVICE_ENDPOINT")
	producerCfg.Region = os.Getenv("LOG_SERVICE_REGION")
	producerCfg.AccessKeyID = os.Getenv("LOG_SERVICE_AK")
	producerCfg.AccessKeySecret = os.Getenv("LOG_SERVICE_SK")

	suite.producer = NewProducer(producerCfg)
	suite.producer.Start()
}

func (suite *SDKProducerTestSuite) TearDownTest() {
	suite.producer.Close()
	_, deleteTopicErr := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topic})
	suite.NoError(deleteTopicErr)
	_, deleteProjectErr := suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.project})
	suite.NoError(deleteProjectErr)
}

func TestSDKProducerTestSuite(t *testing.T) {
	suite.Run(t, new(SDKProducerTestSuite))
}

// TestSendLogs: test send logs
func (suite *SDKProducerTestSuite) TestSendLogs() {
	var logs []*pb.LogGroupList
	for i := 0; i < 10; i++ {
		idx := strconv.Itoa(i)

		logs = append(logs, &pb.LogGroupList{
			LogGroups: []*pb.LogGroup{
				{
					Source:   "localhost",
					FileName: "log" + idx,
					Logs: []*pb.Log{
						{
							Contents: []*pb.LogContent{
								{
									Key:   "key-1",
									Value: "test-message" + idx,
								},
								{
									Key:   "key-2",
									Value: idx,
								},
							},
							Time: time.Now().Unix(),
						},
					},
				},
			},
		})
	}

	// update 10 logGroupLists without compression
	for _, logGroupList := range logs {
		for _, logGroup := range logGroupList.LogGroups {
			suite.NoError(suite.producer.SendLogs("", suite.topic, logGroup.Source, logGroup.FileName, logGroup, nil))
		}
	}

	// wait for consumption
	time.Sleep(60 * time.Second)

	// test search logs
	searchRes, err := suite.cli.SearchLogs(&SearchLogsRequest{
		TopicID:   suite.topic,
		Query:     "*",
		StartTime: 1600000000000,
		EndTime:   2600000000000,
		Limit:     100,
	})
	suite.NoError(err)

	suite.Equal(10, searchRes.Count)

	logMap := make(map[string]struct{})
	for _, searchLog := range searchRes.Logs {
		for _, v := range searchLog {
			switch v.(type) {
			case string:
				logMap[v.(string)] = struct{}{}
			}
		}
	}

	suite.Contains(logMap, "localhost")

	for _, logGroupList := range logs {
		for _, logGroup := range logGroupList.LogGroups {
			suite.Contains(logMap, logGroup.FileName)

			for _, log := range logGroup.Logs {
				for _, content := range log.Contents {
					suite.Contains(logMap, content.Value)
				}
			}
		}
	}
}

// first mock putLogs return 429
func (suite *SDKProducerTestSuite) TestSendLogsWithdrawalMechanism() {

	log := &pb.Log{
		Time: time.Now().Unix(),
		Contents: []*pb.LogContent{
			{
				Key:   "key-1",
				Value: "test-message1",
			},
			{
				Key:   "key-2",
				Value: "test-message2",
			},
		},
	}

	// update 1 logGroupLists without compression
	err := suite.producer.SendLog("", suite.topic, "test-source", "test-filename", log, nil)
	if err != nil {
		return
	}

	// wait for consumption
	time.Sleep(600 * time.Second)
}
