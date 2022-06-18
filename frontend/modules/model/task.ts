import * as Type from './type'
import * as Table from '../table'
export class Tasks {
  tasks = {
    all: [] as Task[],
    sortedTasks: [] as Task[],
    listNumArr: [25, 50, 75, 100],
    listIndex: 0,
    pageIndex: 0,
    basicSortKey: 'id' as keyof Task,
    selectAssignee: '',
    selectField: '',
    selectMilestone: '',
    selectVersion: '',
    selectStatus: [] as string[],
    selectedTask: {},
    taskIndex: -1,
  }
  /** ユーザーのタスク情報を格納 */
  insertTasks(tasks: Type.Task[]) {
    this.tasks.all = this.preprocessTasks(tasks);
    this.tasks.sortedTasks = this.tasks.all;
    this.resetSort();
  }

  /** ユーザーのタスク情報をリセット */
  reset() {
    this.tasks = {
      all: [] as Task[],
      sortedTasks: [] as Task[],
      listNumArr: [25, 50, 75, 100],
      listIndex: 0,
      pageIndex: 0,
      basicSortKey: 'id' as keyof Task,
      selectAssignee: '',
      selectField: '',
      selectMilestone: '',
      selectVersion: '',
      selectStatus: [] as string[],
      selectedTask: {},
      taskIndex: -1,
    }
  }

  /** 表示されるタスクの数listNumArrのインデックス変更 */
  changeListIndex(index: number) {
    this.tasks.pageIndex = 0;
    this.tasks.listIndex = index;
  }

  /** タスクをソートキーを元にソート */
  sortBy(cell: Table.Cell) {
    if('sortKey' in cell === false) return;
    const compare = this.selectFunc(cell.sortKey, cell.header.active, cell.name as keyof Task)
    this.tasks.sortedTasks.sort(compare);
  }
  /** タスクを取り込む際のデータ処理 */
  preprocessTasks = (tasks: Type.Task[]): Task[] => {
    const newTasks = []
    for (let index = 0; index < tasks.length; index++) {
      newTasks.push(this.preprocessTask(tasks[index]));
    }
    return newTasks;
  }
  preprocessTask = (task: Type.Task) => {
    const newTask = {} as Task
    newTask.priority = (task?.priority as Type.Priority)?.name;
    newTask.milestone = (task?.milestone as Type.Milestone)?.name;
    newTask.version = (task?.version as Type.Version)?.name;
    newTask.field = (task.field as Type.Field)?.name;
    newTask.status = (task.status as Type.Status)?.name;
    newTask.type = (task.type as Type.Type)?.name;
    newTask.created_at = this.changeTimeFormat(task.created_at);
    newTask.updated_at = this.changeTimeFormat(task.updated_at);
    newTask.start_time = this.changeTimeFormat(task.start_time);
    newTask.deadline = this.changeTimeFormat(task.deadline);
    return {...task, ...newTask};
  }

  /** 選択されたタスクのオブジェクトを保持 */
  selectTask(params: Type.URLParams) {
    this.tasks.selectedTask = {}
    const t0 = performance.now();
    let key = 0;
    if('key' in params) key = Number(params.key);
    for (let index = 0; index < this.tasks.all.length; index++) {
      if(this.tasks.all[index].id == key) {
        this.tasks.taskIndex = index;
        this.tasks.selectedTask = this.tasks.all[index]
        break
      }
    }
    const t1 = performance.now();
    console.log(`Call to doSomething took ${t1 - t0} milliseconds.`);
    if(Object.keys(this.tasks.selectedTask).length == 0) this.tasks.taskIndex = -1;
  }

  /** 新しいタスクの追加 */
  addTask(task: Type.Task) {
    this.tasks.all.push(this.preprocessTask(task));
    this.tasks.sortedTasks = this.tasks.all;
    this.resetSort();
  }

  /** タスクの更新 */
  updateTask(updatedTask: Type.Task) {
    this.tasks.all = this.tasks.all.map((task) => {
      if(task.id == updatedTask.id) {
        const newTask = this.preprocessTask(updatedTask);
        task = {...{}, ...newTask}
        console.log(task)
      }
      return task;
    })
    this.tasks.sortedTasks = [...[], ...this.tasks.all];
    this.resetSort();
    this.selectTask({key: String(updatedTask.id)} as Type.URLParams)
  }
  changeTimeFormat = (time: string) => {
    const dateObj = new Date(time);
    const year = dateObj.getFullYear();
    const month = dateObj.getMonth() + 1;
    const date = dateObj.getDate();
    const dateStr = year + '/' + month + '/' + date;
    return dateStr;
  }

  /** タイプとキーをもとにソートする関数を返す */
  selectFunc(type: number, activeNum: number, key: keyof Task) {
    if(type === 0) {
      if(activeNum === 1) {
        return function(a: Task, b: Task) {
          if (a[key] > b[key]) return -1;
          if (a[key] < b[key]) return 1;
          return 0;
        }
      }
      return (a: Task, b: Task) => {
        if (a[key] < b[key]) return -1;
        if (a[key] > b[key]) return 1;
        return 0;
      }
    }
    if(type === 1) {
      if(activeNum === 1) {
        return (a: Task, b: Task) => {
          const dateA = a[key] as string;
          const dateB = b[key] as string;
          return new Date(dateB).valueOf() - new Date(dateA).valueOf();
        }
      }
      return (a: Task, b: Task) => {
        const dateA = a[key] as string;
        const dateB = b[key] as string;
        return new Date(dateA).valueOf() - new Date(dateB).valueOf();
      }
    }
    if(type === 2) {
      if(activeNum === 1) {
        return (a: Task, b: Task) => {
          if (a[key] < b[key]) return -1;
          if (a[key] > b[key]) return 1;
          return 0;
        }
      }
      return function(a: Task, b: Task) {
        if (a[key] > b[key]) return -1;
        if (a[key] < b[key]) return 1;
        return 0;
      }
    }
    if(type === 3) {
      // key as keyof Type.User
      if(activeNum === 1) {
        return (a: Task, b: Task) => {
          const userA = a[key] as Type.User
          const userB = b[key] as Type.User
          if (userA['name'] < userB['name']) return -1;
          if (userA['name'] > userB['name']) return 1;
          return 0;
        }
      }
      return function(a: Task, b: Task) {
        const userA = a[key] as Type.User
        const userB = b[key] as Type.User
        if (userA['name'] > userB['name']) return -1;
        if (userA['name'] < userB['name']) return 1;
        return 0;
      }
    }
  }

  /** 選択された担当者名を保持 */
  selectAssignee(name: string) {
    this.tasks.pageIndex = 0;
    this.tasks.selectAssignee = name;
  }

  /** 選択されたフィールドを保持 */
  selectField(name: string) {
    this.tasks.pageIndex = 0;
    this.tasks.selectField = name;
  }

  /** 選択されたマイルストーンを保持 */
  selectMilestone(name: string) {
    this.tasks.pageIndex = 0;
    this.tasks.selectMilestone = name;
  }

  /** 選択されたヴァージョンを保持 */
  selectVersion(name: string) {
    this.tasks.pageIndex = 0;
    this.tasks.selectVersion = name;
  }

  /** 選択されたステータスを保持 */
  selectStatus(statusArr: string[]) {
    this.tasks.pageIndex = 0;
    this.tasks.selectStatus = statusArr;
  }

  /** 並び順をリセット */
  resetSort() {
    this.tasks.sortedTasks.sort((a: Task, b: Task) => {
      const dateB = b[this.tasks.basicSortKey] as string;
      const dateA = a[this.tasks.basicSortKey] as string;
      return new Date(dateB).valueOf() - new Date(dateA).valueOf();
    });
  }

  /** プロジェクトごとにタスクの絞り込み条件をセッションストレージに格納 */
  storeCondition = (v: {}) => {
    const key = location.host + window.$nuxt.$route.params.id;
    let item = sessionStorage.getItem(key);
    let newItem = v;
    if(item) {
      newItem = {...JSON.parse(item), ...v};
    }
    sessionStorage.setItem(key, JSON.stringify(newItem));
  }

  /** タスクの表示数 */
  get listNum(): number {
    return this.tasks.listNumArr[this.tasks.listIndex];
  }

  /** タスクのページ数 */
  get totalPageNum(): number {
    return Math.floor(this.filterTasks.length / this.listNum);
  }

  /** 現在のページに表示されているタスクの最初の番数 */
  get displayedFirstTaskNum(): number {
    const firstList = this.tasks.pageIndex * this.listNum + 1;
    return firstList;
  }

  /** 現在のページに表示されているタスクの最後の番数 */
  get displayedLastTaskNum(): number {
    let lastList = (this.tasks.pageIndex + 1) * this.listNum;
    if(this.filterTasks.length < lastList) lastList = this.filterTasks.length;
    return lastList;
  }
  get currentDisplayedTasksNum(): string {
    return this.displayedFirstTaskNum + ' - ' + this.displayedLastTaskNum;
  }

  /** 担当者、フィールド、マイルストーン、ヴァージョン、ステータスによるタスクのフィルタリング */
  get filterTasks() {
    return this.tasks.sortedTasks.filter((task) => {
      if(!this.tasks.selectStatus.includes('完了') && task.status === '完了') return;
      const assignee = !this.tasks.selectAssignee ?? task.assignee.name == this.tasks.selectAssignee;
      const field =  !this.tasks.selectField ?? task.field == this.tasks.selectField;
      const milestone = !this.tasks.selectMilestone ?? task.milestone == this.tasks.selectMilestone;
      const version = !this.tasks.selectVersion ?? task.version == this.tasks.selectVersion;
      const status = this.tasks.selectStatus.length == 0 ?? this.tasks.selectStatus.includes(task.status);
      return assignee && field && milestone && version && status;
    })
  }

  /** ページによるタスクのフィルタリング */
  get main () {
    return this.filterTasks.filter((_, index) => {
      const start = this.tasks.pageIndex * this.listNum - 1;
      const end = (this.tasks.pageIndex + 1) * this.listNum;
      return start < index && index < end;
    })
  }

  /** 選択されたタスクの情報 */
  get selectedTask () {
    return this.tasks.selectedTask
  }
}

interface Task {
  id: number
  title: string
  assignee: Type.User
  comments: Type.Comment[]
  detail: string
  key: string
  parent: number
  priority: string
  milestone: string
  field: string
  version: string
  status: string
  type: string
  start_time: string
  deadline: string
  estimated_time: number
  actual_time: number
  created_at: string
  updated_at: string
  comment: Type.Comment[]
}