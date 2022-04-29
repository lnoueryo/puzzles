import * as lib from './store/type'
export class MainUser {
  private mainUser = {
    user: {} as lib.User,
    organization: {} as lib.OrganizationAuthority,
    projects: [] as lib.Project[],
    selectedProject: {},
    projectIndex: 0,
  }
  // private _user = {}
  // private _organization = {}
  // private _projects: lib.ProjectAuthority[] = []
  // public _selectedProject = {}
  // private _projectIndex: number = 0;
  insertUser(user: lib.MainUserInfo) {
    this.preprocessUser(user);
  }
  init() {
    this.mainUser = {
      user: {} as lib.User,
      organization: {} as lib.OrganizationAuthority,
      projects: [] as lib.Project[],
      selectedProject: {},
      projectIndex: 0,
    }
  }
  preprocessUser(user: lib.MainUserInfo) {
    this.mainUser = {...this.mainUser, ...user}
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
  get selectedProject() {
    return this.mainUser.selectedProject as lib.ProjectAuthority;
  }
  get projectIndex() {
    return this.mainUser.projectIndex;
  }
  selectProject(params: lib.Params) {
    this.mainUser.projectIndex = 0;
    let id = 0;
    if(Object.keys(params).length != 0) id = Number(params.id);
    this.mainUser.selectedProject = this.mainUser.projects.find((project, index) => {
      this.mainUser.projectIndex = index;
      return project.id == id
    }) ?? {}
    if(Object.keys(this.mainUser.selectedProject).length == 0) this.mainUser.projectIndex = -1;
  }
  createProject(project: lib.Project) {
    this.mainUser.projects.unshift(project)
  }
  updateProject(project: lib.ProjectAuthority) {
    this.mainUser.projects = this.mainUser.projects.map((project) => {
      if(project.id == project.id) {
        project = {...project, ...project}
      }
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
