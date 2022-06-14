import * as Type from './type'
/** ユーザーの情報を管理するオブジェクト */
export class MainUser {
  private mainUser = {
    user: {} as Type.User,
    organizationAuthority: {} as Type.OrganizationAuthority,
    projects: [] as Type.Project[],
    selectedProject: {},
    projectAuthority: {} as Type.ProjectAuthority,
    projectIndex: 0,
    projectID: 0,
    userID: 0,
  }
  insertUser(user: Type.MainUserInfo) {
    this.preprocessUser(user);
  }
  init() {
    this.mainUser = {
      user: {} as Type.User,
      organizationAuthority: {} as Type.OrganizationAuthority,
      projects: [] as Type.Project[],
      selectedProject: {},
      projectAuthority: {} as Type.ProjectAuthority,
      projectIndex: 0,
      projectID: 0,
      userID: 0,
    }
  }
  /** 最初に受けとったデータの前処理 */
  preprocessUser(user: Type.MainUserInfo) {
    this.mainUser = {...this.mainUser, ...user}
    if(!this.mainUser.projects) return this.mainUser.projects = []
    this.mainUser.projects = this.mainUser.projects.sort((a: Type.Project, b: Type.Project) => {
      return new Date(b.created_at).valueOf() - new Date(a.created_at).valueOf();
    });
  }
  /** ユーザーのプロフィール情報 */
  get user() {
    return this.mainUser.user;
  }
  /** ユーザーの組織情報 */
  get organizationAuthority() {
    return this.mainUser.organizationAuthority;
  }
  /** ユーザーのプロジェクト情報 */
  get projects() {
    return this.mainUser.projects;
  }
  /**　プロジェクトをスライドに表示 */
  get projectSlides() {
    const divideConst = 3;
    const projectSlides: Type.Project[][] = []
    this.projects.forEach((project: Type.Project, index: number) => {
      const arrayIndex = Math.floor(index / divideConst);
      if(index % divideConst == 0) {
        projectSlides.splice(arrayIndex, 0, [project]);
      } else {
        projectSlides[arrayIndex].push(project);
      }
    });
    return projectSlides;
  }
  /** 選択されたプロジェクトの情報 */
  get selectedProject(): Type.Project {
    return this.mainUser.projects.find((project) => {
      return project.id == this.mainUser.projectID;
    }) ?? {} as Type.Project;
  }
  /** 組織のオブジェクトから */
  get selectedUser(): Type.OrganizationAuthority {
    return this.mainUser.organizationAuthority?.organization?.users.find((user) => {
      return user.user_id == this.mainUser.userID;
    }) ?? {} as Type.OrganizationAuthority;
  }
  get projectAuthority() {
    console.log(this.selectedProject.authority_users)
    return this.selectedProject.authority_users.find((authority_user) => {
      return authority_user.user_id == this.user.id;
    }) ?? {} as Type.ProjectAuthority
  }
  get projectIndex() {
    return this.mainUser.projectIndex;
  }
  selectProject(params: Type.Params) {
    // this.mainUser.projectIndex = 0;
    let id = 0;
    if('id' in params) id = Number(params.id);
    this.mainUser.projectID = id;
    this.mainUser.projectIndex = this.mainUser.projects.findIndex((project) => id == project.id)
  }
  selectUser(params: Type.Params) {
    let id = 0;
    if('user_id' in params) id = Number(params.user_id);
    this.mainUser.userID = id;
  }
  createProject(project: Type.Project) {
    this.mainUser.projects.unshift(project)
  }
  updateProject(updatedProject: Type.Project) {
    this.mainUser.projects = this.mainUser.projects.map((project) => {
      if(project.id == updatedProject.id) {
        console.log(updatedProject)
        project = {...project, ...updatedProject}
      }
      return project;
    })
  }
  createProjectAuthority(createdProjectAuthority: Type.ProjectAuthority) {
    this.mainUser.projects.forEach((project: Type.Project) => {
      if(project.id != createdProjectAuthority.project_id) return;
      project.authority_users.push(createdProjectAuthority)
    })
  }
  updateProjectAuthority(updateProjectAuthority: Type.ProjectAuthority) {
    this.mainUser.projects = this.mainUser.projects.map((project: Type.Project) => {
      if(project.id != updateProjectAuthority.project_id) return project;
      project.authority_users = project.authority_users.map((user) => {
        if(updateProjectAuthority.user_id != user.user_id) return user;
        user = {...user, ...updateProjectAuthority}
        return user;
      })
      console.log(project.authority_users[1])
      return project;
    })
  }
  deleteProjectAuthority(deleteProjectAuthority: Type.ProjectAuthority) {
    this.mainUser.projects = this.mainUser.projects.map((project: Type.Project) => {
      if(project.id != deleteProjectAuthority.project_id) return project;
      project.authority_users = project.authority_users.filter((user) => {
        return deleteProjectAuthority.user_id != user.user_id;
      })
      return project;
    })
  }
  updateUser(user: Type.User) {
    this.mainUser.user = {...this.mainUser.user, ...user}
  }
}

interface User {
  id: number,
  name: string,
  age: number,
  sex: string,
  email: string,
  address: string,
  image: string,
  description: string,
  organization: string
  authority: string
}
