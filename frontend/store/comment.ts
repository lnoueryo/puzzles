import { GetterTree, ActionTree, MutationTree } from 'vuex'
import * as model from '~/modules/model';
import { Comments } from '~/modules/model/comment';
import { RootState } from '~/store'

const initialState = () => {
  return {
    comments: new Comments(),
    content: '',
    editMode: false,
    selectedComment: {id: 0, index: 0, parent: 0},
  }
}

export const state = initialState()

export type TaskState = ReturnType<typeof initialState>

export const getters: GetterTree<TaskState, RootState> = {
  comments: state => state.comments.comments.all,
  content: state => state.content,
  editMode: state => state.editMode,
  selectedComment: state => state.selectedComment
}

export const mutations: MutationTree<TaskState> = {
  reset: (state) => Object.assign(state, initialState()),
  comments: (state, comments: model.Comment[]) => state.comments.insertComments(comments),
  content: (state, content: string) => state.content = content,
  editMode: (state, editMode: boolean) => state.editMode = editMode,
  addComment: (state, comment: model.Comment) => state.comments.addComment(comment),
  updateComment: (state, comment: model.Comment) => state.comments.updateComment(comment),
  deleteComment: (state, id) => state.comments.deleteComment(id),
  selectComment: (state, comment) => state.selectedComment = comment,
}

export const actions: ActionTree<TaskState, RootState> = {
  reset({commit}) {
    commit('reset');
  },
  editMode({commit}, editMode) {
    commit('editMode', editMode)
    if(!editMode) {
      commit('content', '')
      commit('selectComment', {id: 0, index: 0, parent: 0})
    }
  },
  async getComments({commit, dispatch}, id: number) {
    console.log('Get Comment')
    return new Promise(async(resolve, reject) => {
      try {
        commit('reset')
        const response = await this.$axios.get('/api/comment/show', {
          params: {id: id}
        })
        resolve(response);
        commit('comments', response.data);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async createComment({commit, rootGetters}, form) {
    return new Promise(async(resolve, reject) => {
      try {
        const t0 = performance.now();
        const response = await this.$axios.post('/api/comment/create', form);
        response.data.user = rootGetters.user;
        response.data.replies = [];
        commit('addComment', response.data);
        const t1 = performance.now();
        console.log(`Call to doSomething took ${t1 - t0} milliseconds.`);
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async updateComment({commit, getters}, comment) {
    return new Promise(async(resolve, reject) => {
      const newComment = {...getters.selectedComment}
      newComment.content = getters.content
      console.log(newComment)
      try {
        const response = await this.$axios.put('/api/comment/update', newComment);
        commit('updateComment', response.data);
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  },
  async deleteComment({commit}, comment) {
    return new Promise(async(resolve, reject) => {
      try {
        const commentIDs = [comment.id]
        const treeComments = (comments: model.Comment[]) => {
          comments.forEach((comment) => {
            commentIDs.push(comment.id)
            console.log(commentIDs)
            if(comment.replies?.length != 0) {
              treeComments(comment.replies)
            }
          });
        }
        if(comment.replies?.length != 0) {
          treeComments(comment.replies)
        }
        const response = await this.$axios.delete('/api/comment/delete', {params: {id: commentIDs}});
        commit('deleteComment', comment.id);
        commit('selectComment', {id: 0, index: 0, parent: 0})
        resolve(response);
      } catch (error: any) {
        console.log(error);
        reject(error.response);
      }
    })
  }
}

