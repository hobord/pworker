# PWorker


Worker pool module to run parallel tasks

You sould implement the ```Task``` interface.

```
type Task interface {
	Run()
}
```

You can make simple or buffered Task channel.

Then You can delegate the task to worker pool with method or just send into the
channel.

```
    tasksCH := make(chan Task)
    wp := NewWorkerPool(tasksCH, countOfWorkers)
    
		wp.AddTask(task)
    // or
    taskCH <- task
```

You can add or remove workers from the pool:
```
    wp.AddWorkers(8)
    wp.RemoveWorkers(8)
```

You can get count of workers from the pool:
```
    wp.WorkersCount()
```


You should stop the pool for graceful shutdown.

```
    wp.Stop()
```
