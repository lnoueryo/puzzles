import * as lib from './type'
export class Comments {
    comments = {
    all: [] as lib.Comment[],
    sortedTasks: [] as Comment[],
    listNumArr: [25, 50, 75, 100],
    listIndex: 0,
    pageIndex: 0,
    basicSortKey: 'id' as keyof Comment,
    selectAssignee: '',
    selectField: '',
    selectStatus: [] as string[],
    selectedTask: {},
    taskIndex: -1,
  }
  insertTasks(comments: lib.Comment[]) {
    this.comments.all = comments;
  }
}
