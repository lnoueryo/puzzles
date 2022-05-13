/// <reference types="cypress" />
const homePage = {
  organization: {
    name: '+base',
    description: '看護師を経験し、すべての医療従事者に「心のコップを満たす習慣」を広めようと思った',
    users: [
      {name: '井上領', email: 'popo62520908@gmail.com', authority: '管理者'},
      {name: '君塚 幸介', email: 'kosukekimizuka@example.org', authority: '一般'},
      {name: '木村 佳恵', email: 'yoshie_kimura@example.jp', authority: '一般'},
      {name: '細田 一志', email: 'hosoda105@example.com', authority: '一般'},
      {name: '長井 敬一', email: 'nagaikeiichi@example.org', authority: '一般'},
      {name: '徳村 直哉', email: 'tokumuranaoya@example.org', authority: '一般'},
      {name: '小田 尚', email: 'takashioda@example.com', authority: '一般'},
      {name: '松田 真史', email: 'matsuda_masashi@example.ne.jp', authority: '一般'},
      {name: '野中 健作', email: 'kensakunonaka@example.co.jp', authority: '一般'},
      {name: '緒方 孝', email: 'ogata_takashi@example.com', authority: '一般'},
      {name: '有馬 真理', email: 'arima229@example.ne.jp', authority: '一般'},
      {name: '水野 隆一', email: 'ryuichimizuno@example.co.jp', authority: '一般'},
      {name: '鴨 正哉', email: 'kamomasaya@example.co.jp', authority: '一般'},
      {name: '谷 康司', email: 'tani722@example.ne.jp', authority: '一般'},
      {name: '村上 良平', email: 'ryoheimurakami@example.co.jp', authority: '一般'},
      {name: '渡邉 舞', email: 'mai_watanabe@example.com', authority: '一般'},
      {name: '竹原 友香', email: 'takehara73@example.com', authority: '一般'},
      {name: '大黒 遼太', email: 'daikoku_106@example.ne.jp', authority: '一般'},
      {name: '櫻井 淳一郎', email: 'junichirosakurai@example.co.jp', authority: '一般'},
      {name: '児玉 大', email: 'dai_kodama@example.com', authority: '一般'},
      {name: '上原 良和', email: 'yoshikazu_uehara@example.co.jp', authority: '一般'},
      {name: '井上領', email: 'afsp0908@gmail.com', authority: '一般'},
    ],
    projects: [
      {
        title: 'baseball',
        users: [
          {name: '井上領', email: 'popo62520908@gmail.com', authority: '管理者'},
          {name: '君塚 幸介', email: 'kosukekimizuka@example.org', authority: '管理者'},
          {name: '木村 佳恵', email: 'yoshie_kimura@example.jp', authority: '一般'},
          {name: '細田 一志', email: 'hosoda105@example.com', authority: '一般'},
          {name: '長井 敬一', email: 'nagaikeiichi@example.org', authority: '一般'},
          {name: '小田 尚', email: 'takashioda@example.com', authority: '一般'},
          {name: '松田 真史', email: 'matsuda_masashi@example.ne.jp', authority: '一般'},
          {name: '野中 健作', email: 'kensakunonaka@example.co.jp', authority: '一般'},
          {name: '有馬 真理', email: 'arima229@example.ne.jp', authority: '一般'},
          {name: '鴨 正哉', email: 'kamomasaya@example.co.jp', authority: '一般'},
          {name: '谷 康司', email: 'tani722@example.ne.jp', authority: '一般'},
          {name: '村上 良平', email: 'ryoheimurakami@example.co.jp', authority: '一般'},
          {name: '児玉 大', email: 'dai_kodama@example.com', authority: '一般'},
          {name: '井上領', email: 'afsp0908@gmail.com', authority: '一般'},
        ]
      },
      {
        title: '性格診断',
        users: [
          {name: '井上領', email: 'popo62520908@gmail.com', authority: '管理者'},
          {name: '木村 佳恵', email: 'yoshie_kimura@example.jp', authority: '一般'},
          {name: '長井 敬一', email: 'nagaikeiichi@example.org', authority: '一般'},
          {name: '松田 真史', email: 'matsuda_masashi@example.ne.jp', authority: '一般'},
          {name: '野中 健作', email: 'kensakunonaka@example.co.jp', authority: '一般'},
          {name: '有馬 真理', email: 'arima229@example.ne.jp', authority: '一般'},
          {name: '谷 康司', email: 'tani722@example.ne.jp', authority: '一般'},
          {name: '村上 良平', email: 'ryoheimurakami@example.co.jp', authority: '一般'},
          {name: '渡邉 舞', email: 'mai_watanabe@example.com', authority: '一般'},
          {name: '竹原 友香', email: 'takehara73@example.com', authority: '一般'},
          {name: '大黒 遼太', email: 'daikoku_106@example.ne.jp', authority: '一般'},
          {name: '櫻井 淳一郎', email: 'junichirosakurai@example.co.jp', authority: '一般'},
          {name: '上原 良和', email: 'yoshikazu_uehara@example.co.jp', authority: '一般'},
          {name: '君塚 幸介', email: 'kosukekimizuka@example.org', authority: '管理者'},
        ]
      },
      {
        title: 'insect',
        users: [
          {name: '井上領', email: 'popo62520908@gmail.com', authority: '管理者'},
          {name: '君塚 幸介', email: 'kosukekimizuka@example.org', authority: '一般'},
          {name: '細田 一志', email: 'hosoda105@example.com', authority: '一般'},
          {name: '徳村 直哉', email: 'tokumuranaoya@example.org', authority: '一般'},
          {name: '小田 尚', email: 'takashioda@example.com', authority: '一般'},
          {name: '松田 真史', email: 'matsuda_masashi@example.ne.jp', authority: '一般'},
          {name: '緒方 孝', email: 'ogata_takashi@example.com', authority: '一般'},
          {name: '有馬 真理', email: 'arima229@example.ne.jp', authority: '一般'},
          {name: '鴨 正哉', email: 'kamomasaya@example.co.jp', authority: '一般'},
          {name: '渡邉 舞', email: 'mai_watanabe@example.com', authority: '一般'},
          {name: '竹原 友香', email: 'takehara73@example.com', authority: '一般'},
          {name: '大黒 遼太', email: 'daikoku_106@example.ne.jp', authority: '一般'},
          {name: '児玉 大', email: 'dai_kodama@example.com', authority: '一般'},
          {name: '上原 良和', email: 'yoshikazu_uehara@example.co.jp', authority: '一般'},
        ]
      },
    ]
  }
}
const myProfilePage = {
  mainUser: {
    name: '井上領',
    items: [
      {title: '自己紹介', text: 'テストです'},
      {title: '年齢', text: '31'},
      {title: '性別', text: '男'},
      {title: 'メールアドレス', text: 'popo62520908@gmail.com'},
      {title: '住所', text: '東京都世田谷区松原1-43-14'},
      {title: '組織参加日', text: '2022/4/24'},
    ]
  }
}

const layout = {
  logout: () => {
    cy.get('#profile-list-button', {timeout: 5000}).click()
    cy.get('#profile-list', {timeout: 5000}).contains('ログアウト').click()
  },
  unselectedProject: [
    'puzzles',
    'プロジェクト名',
    'プロジェクトを選択',
    '名前',
    '井上領',
    '組織名',
    '+base',
  ],
  navigationCheck: (items) => {
    items.forEach((item) => {
      cy.get('#header', {timeout: 5000}).should('contain', item)
    })
  }
}

describe('example to-do app', () => {
  beforeEach(() => {
    cy.login('prygen4fDISDVgSYDjxZ5uICD', 'popo62520908@gmail.com', 'popo0908')
  })

  it('check home page project', () => {
    layout.navigationCheck(unselectedProject)
    cy.get('main header', {timeout: 5000}).should('contain', 'プロジェクト')
    cy.get('main header', {timeout: 5000}).should('contain', '組織の概要')
    cy.get('main #organization-name', {timeout: 5000}).should('contain', homePage.organization.name)
    let mainIndex = 0;
    homePage.organization.projects.forEach((project, index) => {
      cy.get('main .v-card', {timeout: 5000}).eq(index).should('contain', project.title)
      cy.get('main .v-card', {timeout: 5000}).eq(index).contains('+10').click()
      cy.get('.v-dialog header', {timeout: 5000}).should('contain', project.title)
      let mainIndex2 = 0;
      project.users.forEach((user, porjectIndex) => {
        cy.get('.v-dialog .v-list-item', {timeout: 10000}).eq(mainIndex + porjectIndex).should('contain', user.name)
        cy.get('.v-dialog .v-list-item', {timeout: 5000}).eq(mainIndex + porjectIndex).should('contain', user.email)
        cy.get('.v-dialog .v-list-item', {timeout: 5000}).eq(mainIndex + porjectIndex).should('contain', user.authority)
        mainIndex2 += 1;
      })
      mainIndex += mainIndex2;
      cy.get('.v-dialog', {timeout: 5000}).eq(index).find('header button').click()
    })
    layout.logout()
  })

  it('check home page summary', () => {
    layout.navigationCheck(unselectedProject)
    cy.get('main header', {timeout: 5000}).should('contain', 'プロジェクト')
    cy.get('main header', {timeout: 5000}).should('contain', '組織の概要').click()
    cy.get('#tab-2 .row', {timeout: 5000}).eq(0).should('contain', homePage.organization.name)
    cy.get('#tab-2 .row', {timeout: 5000}).eq(0).contains('編集').should('be.exist')
    cy.get('#tab-2 .row', {timeout: 5000}).eq(3).should('contain', '概要')
    cy.get('#tab-2 .row', {timeout: 5000}).eq(4).should('contain', homePage.organization.description)
    cy.get('#tab-2 .row', {timeout: 5000}).eq(5).should('contain', 'ユーザー')
    cy.get('#tab-2 .row', {timeout: 5000}).eq(5).contains('メンバー追加').should('be.exist')
    homePage.organization.users.forEach((user, index) => {
      cy.get('main #tab-2 .row', {timeout: 10000}).eq(5).find('.v-list-item').eq(index).should('contain', user.name)
      cy.get('main #tab-2 .row', {timeout: 5000}).eq(5).find('.v-list-item').eq(index).should('contain', user.email)
      cy.get('main #tab-2 .row', {timeout: 5000}).eq(5).find('.v-list-item').eq(index).should('contain', user.authority)
      if(index == 0) {
        cy.get('main #tab-2 .row', {timeout: 5000}).eq(5).find('.v-list-item').eq(index).find('button').click()
        cy.get('.user-options', {timeout: 5000}).eq(index).find('.v-list-item').eq(0).should('have.class', 'v-list-item--disabled')
        cy.get('.user-options', {timeout: 5000}).eq(index).find('.v-list-item').eq(1).should('have.class', 'v-list-item--disabled')
      }
    })
    layout.logout()
  })

  it.only('check profile page', () => {
    cy.get('#profile-list-button', {timeout: 5000}).click()
    cy.get('#profile-list', {timeout: 5000}).contains('プロフィール').click()
    layout.navigationCheck(unselectedProject)
    cy.get('#edit-profile', {timeout: 5000}).should('be.exist')
    cy.get('#user-name', {timeout: 5000}).contains(myProfilePage.mainUser.name)
    myProfilePage.mainUser.items.forEach((item, index) => {
      cy.get('.v-list-item', {timeout: 5000}).eq(index).contains(item.title)
      cy.get('.v-list-item', {timeout: 5000}).eq(index).contains(item.text)
    })
    layout.logout()
  })

  it('can add new todo items', () => {
    // We'll store our item text in a variable so we can reuse it
    const newItem = 'Feed the cat'

    // Let's get the input element and use the `type` command to
    // input our new list item. After typing the content of our item,
    // we need to type the enter key as well in order to submit the input.
    // This input has a data-test attribute so we'll use that to select the
    // element in accordance with best practices:
    // https://on.cypress.io/selecting-elements
    cy.get('[data-test=new-todo]').type(`${newItem}{enter}`)

    // Now that we've typed our new item, let's check that it actually was added to the list.
    // Since it's the newest item, it should exist as the last element in the list.
    // In addition, with the two default items, we should have a total of 3 elements in the list.
    // Since assertions yield the element that was asserted on,
    // we can chain both of these assertions together into a single statement.
    cy.get('.todo-list li')
      .should('have.length', 3)
      .last()
      .should('have.text', newItem)
  })

  it('can check off an item as completed', () => {
    // In addition to using the `get` command to get an element by selector,
    // we can also use the `contains` command to get an element by its contents.
    // However, this will yield the <label>, which is lowest-level element that contains the text.
    // In order to check the item, we'll find the <input> element for this <label>
    // by traversing up the dom to the parent element. From there, we can `find`
    // the child checkbox <input> element and use the `check` command to check it.
    cy.contains('Pay electric bill')
      .parent()
      .find('input[type=checkbox]')
      .check()

    // Now that we've checked the button, we can go ahead and make sure
    // that the list element is now marked as completed.
    // Again we'll use `contains` to find the <label> element and then use the `parents` command
    // to traverse multiple levels up the dom until we find the corresponding <li> element.
    // Once we get that element, we can assert that it has the completed class.
    cy.contains('Pay electric bill')
      .parents('li')
      .should('have.class', 'completed')
  })

  context('with a checked task', () => {
    beforeEach(() => {
      // We'll take the command we used above to check off an element
      // Since we want to perform multiple tests that start with checking
      // one element, we put it in the beforeEach hook
      // so that it runs at the start of every test.
      cy.contains('Pay electric bill')
        .parent()
        .find('input[type=checkbox]')
        .check()
    })

    it('can filter for uncompleted tasks', () => {
      // We'll click on the "active" button in order to
      // display only incomplete items
      cy.contains('Active').click()

      // After filtering, we can assert that there is only the one
      // incomplete item in the list.
      cy.get('.todo-list li')
        .should('have.length', 1)
        .first()
        .should('have.text', 'Walk the dog')

      // For good measure, let's also assert that the task we checked off
      // does not exist on the page.
      cy.contains('Pay electric bill').should('not.exist')
    })

    it('can filter for completed tasks', () => {
      // We can perform similar steps as the test above to ensure
      // that only completed tasks are shown
      cy.contains('Completed').click()

      cy.get('.todo-list li')
        .should('have.length', 1)
        .first()
        .should('have.text', 'Pay electric bill')

      cy.contains('Walk the dog').should('not.exist')
    })

    it('can delete all completed tasks', () => {
      // First, let's click the "Clear completed" button
      // `contains` is actually serving two purposes here.
      // First, it's ensuring that the button exists within the dom.
      // This button only appears when at least one task is checked
      // so this command is implicitly verifying that it does exist.
      // Second, it selects the button so we can click it.
      cy.contains('Clear completed').click()

      // Then we can make sure that there is only one element
      // in the list and our element does not exist
      cy.get('.todo-list li')
        .should('have.length', 1)
        .should('not.have.text', 'Pay electric bill')

      // Finally, make sure that the clear button no longer exists.
      cy.contains('Clear completed').should('not.exist')
    })
  })
})
