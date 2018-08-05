package ns

import (
	"encoding/json"
	"path/filepath"
	"strings"

	etcd "github.com/coreos/etcd/client"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	taskPrefix            = "/agent"
	taskSuffix            = "task"
	taskInfoSuffix        = "request"
	taskServiceTypeSuffix = "serviceType"
	taskTypeSuffix        = "type"
	taskParamSuffix       = "param"
	taskStateSuffix       = "state"
)

func (r *EtcdRegistry) Tasks(agentNode string) (tasks []string, err error) {
	key := r.prefixed(taskPrefix, agentNode, taskSuffix)
	opts := &etcd.GetOptions{
		Sort:      false,
		Recursive: false,
	}

	resp, err := r.kAPI.Get(context.Background(), key, opts)
	if err != nil {
		log.Errorf("Get key: %s, error: %v", key, err)
		if isEtcdError(err, etcd.ErrorCodeKeyNotFound) {
			err = nil
		}
		return
	}

	for _, node := range resp.Node.Nodes {
		tasks = append(tasks, node.Key)
	}

	return
}

func (r *EtcdRegistry) TodoTasks(agentNode string) (tasks []string, err error) {
	key := r.prefixed(taskPrefix, agentNode, taskSuffix)
	opts := &etcd.GetOptions{
		Sort:      false,
		Recursive: false,
	}

	resp, err := r.kAPI.Get(context.Background(), key, opts)
	if err != nil {
		if isEtcdError(err, etcd.ErrorCodeKeyNotFound) {
			err = nil
		} else {
			log.Errorf("Get key: %s, error: %v", key, err)
		}
		return
	}

	for _, node := range resp.Node.Nodes {
		state, err := r.GetTaskState(node.Key)
		if err != nil {
			if isEtcdError(err, etcd.ErrorCodeKeyNotFound) {
				tasks = append(tasks, node.Key)
			} else {
				log.Errorf("Get state: %s, error: %v", node.Key, err)
			}
			continue
		}
		if strings.EqualFold(state, "awaiting") {
			tasks = append(tasks, node.Key)
		}
	}

	return
}

func (r *EtcdRegistry) RunningTasks(agentNode string) (tasks []string, err error) {
	key := r.prefixed(taskPrefix, agentNode, taskSuffix)
	opts := &etcd.GetOptions{
		Sort:      false,
		Recursive: false,
	}

	resp, err := r.kAPI.Get(context.Background(), key, opts)
	if err != nil {
		if isEtcdError(err, etcd.ErrorCodeKeyNotFound) {
			err = nil
		} else {
			log.Errorf("Get key: %s, error: %v", key, err)
		}
		return
	}

	for _, node := range resp.Node.Nodes {
		state, err := r.GetTaskState(node.Key)
		if err != nil {
			if isEtcdError(err, etcd.ErrorCodeKeyNotFound) {
				continue
			} else {
				log.Errorf("Get state: %s, error: %v", node.Key, err)
			}
			continue
		}
		if strings.EqualFold(state, "running") {
			tasks = append(tasks, node.Key)
		}
	}

	return
}

func (r *EtcdRegistry) GetTaskInfo(taskPath string) (ret_info string, idx uint64, err error) {
	key := r.suffixed(taskPath, taskInfoSuffix)
	opts := &etcd.GetOptions{
		Sort:      false,
		Recursive: true,
	}

	resp, err := r.kAPI.Get(context.Background(), key, opts)
	if err != nil {
		log.Errorf("Get key: %s, error: %v", key, err)
		return
	}

	idx = resp.Index
	err = json.Unmarshal([]byte(resp.Node.Value), ret_info)
	if err != nil {
		log.Errorf("Json decode TaskRequestStruct error: %v", resp.Node.Value)
	}

	return
}

func (r *EtcdRegistry) GetTaskServiceType(taskPath string) (task_service_type string, idx uint64, err error) {
	key := r.suffixed(taskPath, taskServiceTypeSuffix)
	opts := &etcd.GetOptions{
		Sort:      false,
		Recursive: true,
	}

	resp, err := r.kAPI.Get(context.Background(), key, opts)
	if err != nil {
		log.Errorf("Get key: %s, error: %v", key, err)
		return
	}

	idx = resp.Index
	task_service_type = resp.Node.Value
	return
}

func (r *EtcdRegistry) GetTaskId(taskPath string) (task_id string, idx uint64, err error) {
	task_id = filepath.Base(taskPath)
	return
}

func (r *EtcdRegistry) GetTaskType(taskPath string) (task_type string, idx uint64, err error) {
	key := r.suffixed(taskPath, taskTypeSuffix)
	opts := &etcd.GetOptions{
		Sort:      false,
		Recursive: true,
	}

	resp, err := r.kAPI.Get(context.Background(), key, opts)
	if err != nil {
		log.Errorf("Get key: %s, error: %v", key, err)
		return
	}

	idx = resp.Index
	task_type = resp.Node.Value
	return
}

func (r *EtcdRegistry) GetTaskParam(taskPath string) (task_params []byte, idx uint64, err error) {
	key := r.suffixed(taskPath, taskParamSuffix)
	opts := &etcd.GetOptions{
		Sort:      false,
		Recursive: true,
	}

	resp, err := r.kAPI.Get(context.Background(), key, opts)
	if err != nil {
		log.Errorf("Get key: %s, error: %v", key, err)
		return
	}

	idx = resp.Index
	task_params = []byte(resp.Node.Value)
	return
}

func (r *EtcdRegistry) GetTaskState(taskPath string) (ret_state string, err error) {
	key := r.suffixed(taskPath, taskStateSuffix)
	opts := &etcd.GetOptions{
		Sort:      false,
		Recursive: false,
	}

	resp, err := r.kAPI.Get(context.Background(), key, opts)
	if err != nil {
		log.Errorf("Get key: %s, error: %v", key, err)
		return
	}

	ret_state = resp.Node.Value

	return
}

func (r *EtcdRegistry) SetTaskState(taskPath string, state string) (err error) {
	key := r.suffixed(taskPath, taskStateSuffix)
	opts := &etcd.SetOptions{
		PrevExist: etcd.PrevIgnore,
	}
	//state.PublishTime = utils.GetCurrentTimeStr()

	log.Debugf("Setting task state key: %s, val: %s", key, state)
	_, err = r.kAPI.Set(context.Background(), key, state, opts)
	if err != nil {
		log.Errorf("Set key: %s, val: %v, error: %v", key, state, err)
		return
	}

	return
}

func (r *EtcdRegistry) ClearTaskState(taskPath string) (err error) {
	key := r.suffixed(taskPath, taskStateSuffix)
	opts := &etcd.DeleteOptions{
		Recursive: true,
	}

	_, err = r.kAPI.Delete(context.Background(), key, opts)
	if err != nil {
		log.Errorf("Delete key: %s, error: %s", key, err)
		return
	}
	return
}
