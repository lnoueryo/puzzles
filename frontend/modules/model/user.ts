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
  };
  selectedUserID = 0;

  /** ユーザーの情報を格納 */
  insertUser(user: Type.MainUserInfo) {
    this.preprocessUser(user);
  }

  /** ユーザーの情報をリセット */
  reset() {
    this.mainUser = {
      user: {} as Type.User,
      organizationAuthority: {} as Type.OrganizationAuthority,
      projects: [] as Type.Project[],
      selectedProject: {},
      projectAuthority: {} as Type.ProjectAuthority,
      projectIndex: 0,
      projectID: 0,
    };
    this.selectedUserID = 0;
  }

  /** 最初に受けとったデータの前処理 */
  preprocessUser(user: Type.MainUserInfo) {
    this.mainUser = {...this.mainUser, ...user};
    if(!this.mainUser.projects) return this.mainUser.projects = [];
    this.mainUser.projects = this.mainUser.projects.sort((a: Type.Project, b: Type.Project) => {
      return new Date(b.created_at).valueOf() - new Date(a.created_at).valueOf();
    });
  }

  /** 選択されたプロジェクトのインデックスを保持 */
  selectProject(params: Type.URLParams) {
    let id = 0;
    if('id' in params) id = Number(params.id);
    this.mainUser.projectID = id;
    this.mainUser.projectIndex = this.mainUser.projects.findIndex((project) => id == project.id);
  }

  /** 選択されたユーザーのオブジェクトを保持 */
  selectUser(params: Type.URLParams) {
    let id = 0;
    if('user_id' in params) id = Number(params.user_id);
    this.selectedUserID = id;
  }

  /** 新しいプロジェクトの追加 */
  createProject(project: Type.Project) {
    this.mainUser.projects.unshift(project);
  }

  /** プロジェクトの更新 */
  updateProject(updatedProject: Type.Project) {
    this.mainUser.projects = this.mainUser.projects.map((project) => {
      if(project.id == updatedProject.id) project = {...project, ...updatedProject};
      return project;
    });
  }

  /** プロジェクトのユーザー追加 */
  createProjectAuthority(createdProjectAuthority: Type.ProjectAuthority) {
    this.mainUser.projects.forEach((project: Type.Project) => {
      if(project.id != createdProjectAuthority.project_id) return;
      project.authority_users.push(createdProjectAuthority);
    });
  }

  /** プロジェクトのユーザー権限変更 */
  updateProjectAuthority(updateProjectAuthority: Type.ProjectAuthority) {
    this.mainUser.projects = this.mainUser.projects.map((project: Type.Project) => {
      if(project.id != updateProjectAuthority.project_id) return project;
      project.authority_users = project.authority_users.map((user) => {
        if(updateProjectAuthority.user_id != user.user_id) return user;
        user = {...user, ...updateProjectAuthority}
        return user;
      });
      return project;
    })
  }

  /** プロジェクトのユーザー削除 */
  deleteProjectAuthority(deleteProjectAuthority: Type.ProjectAuthority) {
    this.mainUser.projects = this.mainUser.projects.map((project: Type.Project) => {
      if(project.id != deleteProjectAuthority.project_id) return project;
      project.authority_users = project.authority_users.filter((user) => {
        return deleteProjectAuthority.user_id != user.user_id;
      });
      return project;
    });
  }

  /** ユーザー情報変更 */
  updateUser(user: Type.User) {
    this.mainUser.user = {...this.mainUser.user, ...user};
  }

  /** ユーザーのプロフィール情報 */
  get user(): Type.User {
    return this.mainUser.user;
  }

  /** ユーザーの組織情報 */
  get organizationAuthority(): Type.OrganizationAuthority {
    return this.mainUser.organizationAuthority;
  }

  /** ユーザーのプロジェクト情報 */
  get projects(): Type.Project[] {
    return this.mainUser.projects;
  }

  /**　プロジェクトをスライドに表示 */
  get projectSlides(): Type.Project[][] {
    const divideConst = 3;
    const projectSlides: Type.Project[][] = []
    this.projects.forEach((project: Type.Project, index: number) => {
      const arrayIndex = Math.floor(index / divideConst);
      if(index % divideConst == 0) return projectSlides.splice(arrayIndex, 0, [project]);
      projectSlides[arrayIndex].push(project);
    });
    return projectSlides;
  }

  /** 選択されたプロジェクトの情報 */
  get selectedProject(): Type.Project {
    return this.mainUser.projects.find((project) => {
      return project.id == this.mainUser.projectID;
    }) ?? {} as Type.Project;
  }

  /** ユーザーの組織に属する別ユーザーの情報 */
  get selectedUser(): Type.OrganizationAuthority {
    return this.mainUser.organizationAuthority?.organization?.users.find((user) => {
      return user.user_id == this.selectedUserID;
    }) ?? {} as Type.OrganizationAuthority;
  }

  /** 選択されているプロジェクトに対するユーザーの権限情報 */
  get projectAuthority(): Type.ProjectAuthority {
    return this.selectedProject.authority_users.find((authority_user) => {
      return authority_user.user_id == this.user.id;
    }) ?? {} as Type.ProjectAuthority
  }

  /** 選択されているプロジェクトのインデックス */
  get projectIndex(): number {
    return this.mainUser.projectIndex;
  }
}

