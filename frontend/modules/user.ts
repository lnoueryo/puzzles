import * as lib from './store/type'
export class MainUser {
  private mainUser = {
    user: {} as lib.User,
    organization: {} as lib.OrganizationAuthority,
    projects: [] as lib.Project[],
    selectedProject: {},
    projectAuthority: {} as lib.ProjectAuthority,
    projectIndex: 0,
    projectID: 0,
    userID: 0,
  }
  insertUser(user: lib.MainUserInfo) {
    this.preprocessUser(user);
  }
  init() {
    this.mainUser = {
      user: {} as lib.User,
      organization: {} as lib.OrganizationAuthority,
      projects: [] as lib.Project[],
      selectedProject: {},
      projectAuthority: {} as lib.ProjectAuthority,
      projectIndex: 0,
      projectID: 0,
      userID: 0,
    }
  }
  preprocessUser(user: lib.MainUserInfo) {
    this.mainUser = {...this.mainUser, ...user}
    if(!this.mainUser.projects) return this.mainUser.projects = []
    this.mainUser.projects = this.mainUser.projects.sort((a: lib.Project, b: lib.Project) => {
      return new Date(b.created_at).valueOf() - new Date(a.created_at).valueOf();
    });
  }
  get user() {
    return this.mainUser.user;
  }
  get organization() {
    return this.mainUser.organization;
  }
  get projects() {
    return this.mainUser.projects;
  }
  get projectSlides() {
    const divideConst = 3;
    const projectSlides: lib.Project[][] = []
    this.projects.forEach((project: lib.Project, index: number) => {
      const arrayIndex = Math.floor(index / divideConst);
      if(index % divideConst == 0) {
        projectSlides.splice(arrayIndex, 0, [project]);
      } else {
        projectSlides[arrayIndex].push(project);
      }
    });
    return projectSlides;
  }
  // get selectedProject() {
  //   return this.mainUser.selectedProject as lib.Project;
  // }
  get selectedProject(): lib.Project {
    return this.mainUser.projects.find((project) => {
      return project.id == this.mainUser.projectID;
    }) ?? {} as lib.Project;
  }
  get selectedUser(): lib.OrganizationAuthority {
    return this.mainUser.organization?.organization?.users.find((user) => {
      return user.user_id == this.mainUser.userID;
    }) ?? {} as lib.OrganizationAuthority;
  }
  get projectAuthority() {
    console.log(this.selectedProject.authority_users)
    return this.selectedProject.authority_users.find((authority_user) => {
      return authority_user.user_id == this.user.id;
    }) ?? {} as lib.ProjectAuthority
  }
  get projectIndex() {
    return this.mainUser.projectIndex;
  }
  selectProject(params: lib.Params) {
    // this.mainUser.projectIndex = 0;
    let id = 0;
    if('id' in params) id = Number(params.id);
    this.mainUser.projectID = id;
    this.mainUser.projectIndex = this.mainUser.projects.findIndex((project) => id == project.id)
  }
  selectUser(params: lib.Params) {
    let id = 0;
    if('user_id' in params) id = Number(params.user_id);
    this.mainUser.userID = id;
  }
  createProject(project: lib.Project) {
    this.mainUser.projects.unshift(project)
  }
  updateProject(updatedProject: lib.Project) {
    this.mainUser.projects = this.mainUser.projects.map((project) => {
      if(project.id == updatedProject.id) {
        console.log(updatedProject)
        project = {...project, ...updatedProject}
      }
      return project;
    })
  }
  createProjectAuthority(createdProjectAuthority: lib.ProjectAuthority) {
    this.mainUser.projects.forEach((project: lib.Project) => {
      if(project.id != createdProjectAuthority.project_id) return;
      project.authority_users.push(createdProjectAuthority)
    })
  }
  updateProjectAuthority(updateProjectAuthority: lib.ProjectAuthority) {
    this.mainUser.projects = this.mainUser.projects.map((project: lib.Project) => {
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
  deleteProjectAuthority(deleteProjectAuthority: lib.ProjectAuthority) {
    this.mainUser.projects = this.mainUser.projects.map((project: lib.Project) => {
      if(project.id != deleteProjectAuthority.project_id) return project;
      project.authority_users = project.authority_users.filter((user) => {
        return deleteProjectAuthority.user_id != user.user_id;
      })
      return project;
    })
  }
  updateUser(user: lib.User) {
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
