export interface User {
  id: number,
  name: string,
  age: number,
  sex: string,
  email: string,
  address: string,
  image: string,
  description: string,
  organization: string
  organizations: Organization[]
  projects: ProjectAuthority[]
  authority: string
}

export interface Organization {
  id: number,
  name: string,
  address: string,
  image: string,
  description: string,
  plan: string,
  founded: string,
  credit_card: string,
  expiry_date: string,
}

export interface Project {
  id: number,
  name: string,
  address: string,
  image: string,
  description: string,
  tasks: Task[]
  users: User[]
  fields: Field[]
  milestones: Milestone[]
  authority_users: ProjectAuthority[]
  image_data: string
  created_at: string
  organization_id: string
}

export interface ProjectAuthority {
  id: number
  user_id: number,
  project_id: number,
  auth_id: number,
  type_id: number,
  user: User,
  type: Authority
  active: boolean
  project: Project
  project_users: ProjectAuthority[]
}

export interface Task {
  id: number
  title: string
  assignee: User
  assignee_id: number
  assigner: User
  comments: Comment[]
  detail: string
  key: string
  parent: number
  priority: Priority | string
  priority_id: number
  milestone: Milestone | string
  milestone_id: number
  field: Field | string
  field_id: number
  status: Status | string
  status_id: number
  type: Type | string
  type_id: number
  start_time: string
  deadline: string
  estimated_time: number
  actual_time: number
  created_at: string
  updated_at: string
}


export interface Field {
  id: number
  name: string
}

export interface Milestone {
  id: number
  name: string
}

export interface Status {
  id: number
  name: string
}

export interface Type {
  id: number
  name: string
}

export interface Priority {
  id: number
  name: string
}

export interface OrganizationAuthority {
  type: string
}

export interface Comment {
  type: string
}

export interface Table {
  style: {minWidth: string | number}
  thead: {style: {minWidth: string, backgroundColor: string}},
  tbody: {style: {minWidth: string, overflowX: string, overflowY: string, maxHeight: string}},
  cells: Cell[]
}
export interface Cell {
  name: keyof Task
  header: {title: string, active: number, style: {width: string | number}}
  sortKey: number
}
export interface Authority {
  id: number
  name: string
}

export interface Params {
  id: string
  key: string
}