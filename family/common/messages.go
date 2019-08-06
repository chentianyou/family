package common

import "sync"

type ResponseMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// response error code
const (
	SUCCESS = iota
	REQUEST_DATA_ERROR
	JSON_ERROR
	GET_DATA_ERROR
	UPDATE_DATA_ERROR
	DELETE_DATA_ERROR
	KEY_ID_ERROR
	NAME_DUPLICATE_ERROR
	PERMISSION_DENIED
	AUTHORIZATION_ERROR
	TRAINNODE_ZERO
	NO_GRAPH_IS_RUNNING
	FLOW_TYPE_ERROR
	FLOW_GRAPH_ERROR
	UPDATE_HDFS_URL_FAILED
	GRAPH_IS_STOPPING
	TRAIN_NODE_ERROR
	PREDICT_ERROR
	EXPORT_MODEL_ERROR
	GET_COOKIE_ERROR
	DELETE_CHECK_ERROR
	OTHER_ERROR
)

var responseMsg ResponseMessagePool

var MSG_SUCCESS ResponseMessage
var MSG_REQUEST_DATA_ERROR ResponseMessage
var MSG_JSON_ERROR ResponseMessage
var MSG_GET_DATA_ERROR ResponseMessage
var MSG_UPDATE_DATA_ERROR ResponseMessage
var MSG_DELETE_DATA_ERROR ResponseMessage
var MSG_KEY_ID_ERROR ResponseMessage
var MSG_NAME_DUPLICATE_ERROR ResponseMessage
var MSG_PERMISSION_DENIED ResponseMessage
var MSG_AUTHORIZATION_ERROR ResponseMessage
var MSG_TRAINNODE_ZERO ResponseMessage
var MSG_NO_GRAPH_IS_RUNNING ResponseMessage
var MSG_FLOW_TYPE_ERROR ResponseMessage
var MSG_FLOW_GRAPH_ERROR ResponseMessage
var MSG_UPDATE_HDFS_URL_FAILED ResponseMessage
var MSG_GRAPH_IS_STOPPING ResponseMessage
var MSG_TRAIN_NODE_ERROR ResponseMessage
var MSG_PREDICT_ERROR ResponseMessage
var MSG_EXPORT_MODEL_ERROR ResponseMessage
var MSG_OTHER_ERROR = func(msg string) ResponseMessage {
	return ResponseMessage{OTHER_ERROR, msg}
}
var MSG_GET_COOKIE_ERROR ResponseMessage

func init() {
	responseMsg.messages = map[int]string{
		SUCCESS:                "成功",
		REQUEST_DATA_ERROR:     "获取请求数据失败",
		JSON_ERROR:             "JSON数据错误",
		GET_DATA_ERROR:         "获取数据失败",
		UPDATE_DATA_ERROR:      "更新数据失败",
		DELETE_DATA_ERROR:      "删除数据失败",
		KEY_ID_ERROR:           "关键字ID错误",
		NAME_DUPLICATE_ERROR:   "名称重复",
		PERMISSION_DENIED:      "权限不足",
		AUTHORIZATION_ERROR:    "身份认证错误",
		TRAINNODE_ZERO:         "没有发现训练机器",
		NO_GRAPH_IS_RUNNING:    "没有正在运行的流程图",
		FLOW_TYPE_ERROR:        "非法的流程类型",
		FLOW_GRAPH_ERROR:       "流程图错误",
		UPDATE_HDFS_URL_FAILED: "无法访问HDFS训练数据",
		GRAPH_IS_STOPPING:      "正在停止训练",
		TRAIN_NODE_ERROR:       "训练节点异常",
		PREDICT_ERROR:          "预测错误",
		EXPORT_MODEL_ERROR:     "导出模型错误",
		GET_COOKIE_ERROR:       "获取cookie出错",
	}
	MSG_SUCCESS = ResponseMessageFunc(SUCCESS)
	MSG_REQUEST_DATA_ERROR = ResponseMessageFunc(REQUEST_DATA_ERROR)
	MSG_JSON_ERROR = ResponseMessageFunc(JSON_ERROR)
	MSG_GET_DATA_ERROR = ResponseMessageFunc(GET_DATA_ERROR)
	MSG_UPDATE_DATA_ERROR = ResponseMessageFunc(UPDATE_DATA_ERROR)
	MSG_DELETE_DATA_ERROR = ResponseMessageFunc(DELETE_DATA_ERROR)
	MSG_KEY_ID_ERROR = ResponseMessageFunc(KEY_ID_ERROR)
	MSG_NAME_DUPLICATE_ERROR = ResponseMessageFunc(NAME_DUPLICATE_ERROR)
	MSG_PERMISSION_DENIED = ResponseMessageFunc(PERMISSION_DENIED)
	MSG_AUTHORIZATION_ERROR = ResponseMessageFunc(AUTHORIZATION_ERROR)
	MSG_TRAINNODE_ZERO = ResponseMessageFunc(TRAINNODE_ZERO)
	MSG_NO_GRAPH_IS_RUNNING = ResponseMessageFunc(NO_GRAPH_IS_RUNNING)
	MSG_FLOW_TYPE_ERROR = ResponseMessageFunc(FLOW_TYPE_ERROR)
	MSG_FLOW_GRAPH_ERROR = ResponseMessageFunc(FLOW_GRAPH_ERROR)
	MSG_UPDATE_HDFS_URL_FAILED = ResponseMessageFunc(UPDATE_HDFS_URL_FAILED)
	MSG_GRAPH_IS_STOPPING = ResponseMessageFunc(GRAPH_IS_STOPPING)
	MSG_TRAIN_NODE_ERROR = ResponseMessageFunc(TRAIN_NODE_ERROR)
	MSG_PREDICT_ERROR = ResponseMessageFunc(PREDICT_ERROR)
	MSG_EXPORT_MODEL_ERROR = ResponseMessageFunc(EXPORT_MODEL_ERROR)
	MSG_GET_COOKIE_ERROR = ResponseMessageFunc(GET_COOKIE_ERROR)
}

func (e *ResponseMessagePool) ResponseMessage(code int) ResponseMessage {
	resp := ResponseMessage{}
	resp.Code = code
	e.mutex.RLock()
	resp.Message = e.messages[code]
	e.mutex.RUnlock()
	return resp
}

var ResponseMessageFunc = responseMsg.ResponseMessage

type ResponseMessagePool struct {
	messages map[int]string
	mutex    sync.RWMutex
}
