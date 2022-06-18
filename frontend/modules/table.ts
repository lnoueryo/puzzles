import * as Type from './model/type'

export interface TableData {
  style: {minWidth: string | number}
  thead: {style: {minWidth: string, backgroundColor: string}},
  tbody: {style: {minWidth: string, overflowX: string, overflowY: string, maxHeight: string}},
  cells: Cell[]
}
export interface Cell {
  name: keyof Type.Task
  header: {title: string, active: number, style: {width: string | number}}
  sortKey: number
}

const CELL_WIDTH_A = 80
const CELL_WIDTH_B = 105
const CELL_WIDTH_C = 130
const CELL_WIDTH_D = 160
const CELL_WIDTH_E = 220
const SCROLLBAR_WIDTH = 5

export class Table {
  readonly CELL_WIDTH_A = CELL_WIDTH_A
  readonly CELL_WIDTH_B = CELL_WIDTH_B
  readonly CELL_WIDTH_C = CELL_WIDTH_C
  readonly CELL_WIDTH_D = CELL_WIDTH_D
  readonly CELL_WIDTH_E = CELL_WIDTH_E
  readonly SCROLLBAR_WIDTH = SCROLLBAR_WIDTH
  cells = [
    {name: "key", header: {title: 'タスクキー', active: 0, style: {width: this.CELL_WIDTH_D}}},
    {name: "field", sortKey: 0, header: {title: '分野', active: 0, style: {width: this.CELL_WIDTH_C}}},
    {name: "title", sortKey: 0, header: {title: 'タイトル', active: 0, style: {width: this.CELL_WIDTH_E}}},
    {name: "assignee", sortKey: 3, header: {title: '担当者', active: 0, style: {width: this.CELL_WIDTH_D}}},
    {name: "status", sortKey: 2, header: {title: '状況', active: 0, style: {width: this.CELL_WIDTH_A}}},
    {name: "priority", sortKey: 2, header: {title: '優先', active: 0, style: {width: this.CELL_WIDTH_A}}},
    {name: "milestone", sortKey: 0, header: {title: 'マイルストーン', active: 0, style: {width: this.CELL_WIDTH_D}}},
    {name: "version", sortKey: 0, header: {title: 'バージョン', active: 0, style: {width: this.CELL_WIDTH_D}}},
    {name: "type", sortKey: 0, header: {title: '種類', active: 0, style: {width: this.CELL_WIDTH_A}}},
    {name: 'estimated_time', sortKey: 2, header: {title: '予定', active: 0, style: {width: this.CELL_WIDTH_A}}},
    {name: "deadline", sortKey: 1, header: {title: '期限日', active: 0, style: {width: this.CELL_WIDTH_B}}},
    {name: "assigner", sortKey: 3, header: {title: '作成者', active: 0, style: {width: this.CELL_WIDTH_D}}},
    {name: "created_at", sortKey: 1, header: {title: '作成日', active: 0, style: {width: this.CELL_WIDTH_B}}},
    {name: "updated_at", sortKey: 1, header: {title: '更新日', active: 0, style: {width: this.CELL_WIDTH_B}}},
    {name: "start_time", sortKey: 1, header: {title: '開始日', active: 0, style: {width: this.CELL_WIDTH_B}}},
  ] as Cell[]
  items = {
    style: {minWidth: 0},
    thead: {style: {minWidth: '', backgroundColor: '#295caa'}},
    tbody: {style: {minWidth: '', maxHeight: '500px'}},
    cells: this.cells
  }
  constructor() {
    let tableWidth = 0
    this.items.cells = this.cells.map((cell) => {
      tableWidth += (cell.header.style.width as number);
      cell.header.style.width = cell.header.style.width + 'px'
      return cell;
    })
    this.items.thead.style.minWidth = tableWidth + this.SCROLLBAR_WIDTH + 'px';
    this.items.tbody.style.minWidth = tableWidth + this.SCROLLBAR_WIDTH + 'px';
  }
  resetActive = (cells: Cell[]) => {
    for (let index = 0; index < cells.length; index++) {
      cells[index].header.active = 0;
    }
    return cells
  }
}
