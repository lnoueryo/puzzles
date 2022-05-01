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
  organizations: OrganizationAuthority[]
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
  projects: ProjectAuthority[]
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
  user_id: number
  organization: Organization
}

export interface Comment {
  id: number
  content: string
  user: User
  parent_id: number
  replies: Comment[]
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

export interface Response {
  data: {message: string}
  status: number
}

export interface MainUserInfo {
  user: User
  organization: OrganizationAuthority
  projects: Project[]
}