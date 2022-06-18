import type { User, Organization, Project, Task, Field, Milestone, Version, Status, Type, Priority, OrganizationAuthority, Authority, ProjectAuthority, Comment, URLParams, Response, MainUserInfo } from './type'
import { MainUser } from './user';
import { Tasks } from './task';
export const user = new MainUser()
export const statuses = [
    {id: 1, name: '相談'},
    {id: 2, name: '依頼'},
    {id: 3, name: '再議'},
    {id: 4, name: '未対応'},
    {id: 5, name: '対応中'},
    {id: 6, name: '中断'},
    {id: 7, name: '確認'},
    {id: 8, name: '調整'},
    {id: 9, name: '完了'},
  ]
  
  export const priorities = [
    {id: 1, name: '低'},
    {id: 2, name: '中'},
    {id: 3, name: '高'},
  ]
  
  export const types = [
    {id: 1, name: '追加'},
    {id: 2, name: '変更'},
    {id: 3, name: 'バグ'},
    {id: 4, name: 'その他'},
  ]
  
  export const authorities = [
    {id: 1, name: '管理者'},
    {id: 2, name: '一般'},
  ]
export { User, Organization, Project, Task, Field, Milestone, Version, Status, Type, Priority, OrganizationAuthority, Authority, ProjectAuthority, Comment, URLParams, Response, MainUserInfo }
export { Tasks }




