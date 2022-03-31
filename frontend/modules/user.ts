import * as lib from './store/type'
export class MainUser {
  private mainUser = {
    user: {},
    organization: {},
    projects: [] as lib.ProjectAuthority[],
    selectedProject: {},
    projectIndex: 0,
  }
  // private _user = {}
  // private _organization = {}
  // private _projects: lib.ProjectAuthority[] = []
  // public _selectedProject = {}
  // private _projectIndex: number = 0;
  insertUser(user: lib.User) {
    this.preprocessUser(user);
  }
  init() {
    this.mainUser = {
      user: {},
      organization: {},
      projects: [] as lib.ProjectAuthority[],
      selectedProject: {},
      projectIndex: 0,
    }
  }
  preprocessUser(user: lib.User) {
    this.mainUser.organization = user.organizations[0] as lib.Organization;
    this.mainUser.projects = user.projects;
    this.mainUser.projects = this.mainUser.projects.sort((a: lib.ProjectAuthority, b: lib.ProjectAuthority) => {
      return new Date(b.project.created_at).valueOf() - new Date(a.project.created_at).valueOf();
    });
    user.organizations = [];
    user.projects = [];
    this.mainUser.user = user as User;
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
    const projectSlides: lib.ProjectAuthority[][] = []
    this.projects.forEach((projectAuthority: lib.ProjectAuthority, index: number) => {
      const arrayIndex = Math.floor(index / divideConst);
      if(index % divideConst == 0) {
        projectSlides.splice(arrayIndex, 0, [projectAuthority]);
      } else {
        projectSlides[arrayIndex].push(projectAuthority);
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
      return project.project_id == id
    }) ?? {}
    // console.log(this.mainUser.projects,123)
    // console.log(params,123)
    // console.log(Object.keys(this.mainUser.selectedProject).length == 0)
    if(Object.keys(this.mainUser.selectedProject).length == 0) this.mainUser.projectIndex = -1;
    console.log(this.mainUser.projectIndex)
  }
  createProject(projectAuthority: lib.ProjectAuthority) {
    this.mainUser.projects.unshift(projectAuthority)
  }
  updateProject(projectAuthority: lib.ProjectAuthority) {
    this.mainUser.projects = this.mainUser.projects.map((project) => {
      if(project.project_id == projectAuthority.project_id) {
        project.project = {...project.project, ...projectAuthority.project}
        project = {...project, ...projectAuthority}
      }
      return project;
    })
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
