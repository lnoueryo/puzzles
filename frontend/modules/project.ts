import * as lib from './store/type'
export class SelectProject {
  private _project: Project = {
    id: 0,
    name: '',
    address: '',
    image: '',
    description: '',
    fields: [],
    milestones: [],
    authority_users: []
  }
  insertProject(project: lib.Project) {
    this.preprocessProject(project);
  }
  preprocessProject(project: lib.Project) {
    project.tasks = []
    this._project = project;
  }
  get main() {
    return this._project;
  }
}

interface Project {
  id: number,
  name: string,
  address: string,
  image: string,
  description: string,
  fields: lib.Field[]
  milestones: lib.Milestone[]
  authority_users: lib.ProjectAuthority[]
}