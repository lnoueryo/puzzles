import { Cell, Table } from './type'

export const resetActive = (cells: Cell[]) => {
  for (let index = 0; index < cells.length; index++) {
    cells[index].header.active = 0;
  }
  return cells
}

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

const CELL_WIDTH_A = 80
const CELL_WIDTH_B = 105
const CELL_WIDTH_C = 130
const CELL_WIDTH_D = 160
const CELL_WIDTH_E = 220
const SCROLLBAR_WIDTH = 5

// stringは0, dateは1, numberは2, Userは3、それ以外はtable.tsで型ごとに新たに追加。タスクキーはなし
// 表示順番の変更はここで行う
export const cells = [
  {name: "key", header: {title: 'タスクキー', active: 0, style: {width: CELL_WIDTH_D}}},
  {name: "field", sortKey: 0, header: {title: '分野', active: 0, style: {width: CELL_WIDTH_C}}},
  {name: "title", sortKey: 0, header: {title: 'タイトル', active: 0, style: {width: CELL_WIDTH_E}}},
  {name: "assignee", sortKey: 3, header: {title: '担当者', active: 0, style: {width: CELL_WIDTH_D}}},
  {name: "status", sortKey: 2, header: {title: '状況', active: 0, style: {width: CELL_WIDTH_A}}},
  {name: "priority", sortKey: 2, header: {title: '優先', active: 0, style: {width: CELL_WIDTH_A}}},
  {name: "milestone", sortKey: 0, header: {title: 'マイルストーン', active: 0, style: {width: CELL_WIDTH_D}}},
  {name: "type", sortKey: 0, header: {title: '種類', active: 0, style: {width: CELL_WIDTH_A}}},
  {name: 'estimated_time', sortKey: 2, header: {title: '予定', active: 0, style: {width: CELL_WIDTH_A}}},
  {name: "deadline", sortKey: 1, header: {title: '期限日', active: 0, style: {width: CELL_WIDTH_B}}},
  {name: "assigner", sortKey: 3, header: {title: '作成者', active: 0, style: {width: CELL_WIDTH_D}}},
  {name: "created_at", sortKey: 1, header: {title: '作成日', active: 0, style: {width: CELL_WIDTH_B}}},
  {name: "updated_at", sortKey: 1, header: {title: '更新日', active: 0, style: {width: CELL_WIDTH_B}}},
  {name: "start_time", sortKey: 1, header: {title: '開始日', active: 0, style: {width: CELL_WIDTH_B}}},
] as Cell[]

const table = {
  style: {minWidth: 0},
  thead: {style: {minWidth: '', backgroundColor: '#295caa'}},
  tbody: {style: {minWidth: '', maxHeight: '500px'}},
  cells: cells
} as Table

export const preprocessTable = () => {
  let tableWidth = 0
  const copyCells = JSON.parse(JSON.stringify(cells))
  table.cells = copyCells.map((cell: Cell) => {
    tableWidth += (cell.header.style.width as number);
    cell.header.style.width = cell.header.style.width + 'px'
    return cell;
  })
  table.thead.style.minWidth = tableWidth + SCROLLBAR_WIDTH + 'px';
  table.tbody.style.minWidth = tableWidth + SCROLLBAR_WIDTH + 'px';
  return table
}

export const storeCondition = (v: {}) => {
  const key = location.host + window.$nuxt.$route.params.id;
  let item = sessionStorage.getItem(key);
  let newItem = v;
  if(item) {
    newItem = {...JSON.parse(item), ...v};
  }
  sessionStorage.setItem(key, JSON.stringify(newItem));
}