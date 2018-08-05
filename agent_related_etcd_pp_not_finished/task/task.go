package taskManager

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type TaskManager struct {
	reg         ns.EtcdRegistry
	mach        *ns.Machine
	pkgSite     string
	etcdServers string
}

const stepInterval = time.Second * 10

func NewTaskManager(reg ns.EtcdRegistry, mach *ns.Machine, pkgSite string, etcdServers string) *TaskManager {
	return &TaskManager{
		reg:         reg,
		mach:        mach,
		pkgSite:     pkgSite,
		etcdServers: etcdServers,
	}
}

func (tm *TaskManager) handleTask(task_path, task_id string, task_param []byte, task_type string) (err error) {
	// Do your stuff here..
	return
}

func (tm *TaskManager) ScheduleTask(stop <-chan struct{}) (err error) {
	var tasks []string
	// get task list
	tasks, err = tm.reg.TodoTasks(tm.mach.State().InstanceRegIP)

	if err != nil {
		log.Errorf("Get todo task failed, err: %v", err)
		return
	}
	if log.GetLevel() == log.DebugLevel {
		log.Debugln("Getting tasks:")
		for i, t := range tasks {
			log.Debugf("  %05d: %s", i, t)
		}
	}

	for _, task_path := range tasks {
		select {
		case <-stop:
			log.Infof("    Stop task %s", task_path)
			break
		default:
			var task_param []byte
			task_param, _, err = tm.reg.GetTaskParam(task_path)
			if err != nil {
				log.Errorf("Get task params error, path: %s", task_path)
				continue
			}
			task_type, _, err := tm.reg.GetTaskType(task_path)
			if err != nil {
				log.Errorf("Get task type error, path: %s", task_path)
				continue
			}
			task_id, _, err := tm.reg.GetTaskId(task_path)
			if err != nil {
				log.Errorf("Get task Id error, path: %s", task_path)
			}

			log.Debugf("Task params %s", task_param)

			// 分布式协调中心下放的task处理
			tm.handleTask(task_path, task_id, task_param, task_type)

			tm.mach.Refresh()

			if err != nil {
				log.Print("Setting task: %s state failed", task_path)
			}
			if log.GetLevel() == log.DebugLevel {
				time.Sleep(stepInterval)
			}
		}
	}
	return
}

func (tm *TaskManager) PeriodicSchedule(interval time.Duration, stop <-chan struct{}, mstopc chan struct{}) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-stop:
			log.Print("Halting TaskManager.PeriodicSchedule")
			ticker.Stop()
			close(mstopc)
			return
		case <-ticker.C:
			tm.ScheduleTask(stop)
		}
	}
}
