import * as lib from './type'
export class Comments {
    comments = {
    all: [] as lib.Comment[],
  }
  insertComments(comments: lib.Comment[]) {
    this.comments.all = comments;
  }
  addComment(newComment: lib.Comment) {
    if(newComment.parent_id == 0) {
      this.comments.all.push(newComment);
      return
    }
    const treeComments = (comments: lib.Comment[]) => {
      return comments.map((comment: lib.Comment) => {
        if(comment.id == newComment.parent_id) {
          comment.replies.push(newComment);
          return comment
        }
        if(comment?.replies?.length != 0) comment.replies = treeComments(comment.replies);
        return comment
      })
    }
    this.comments.all = this.comments.all.map((comment: lib.Comment) => {
      if(comment.id == newComment.parent_id) {
        comment.replies.push(newComment);
        return comment
      }
      if(comment?.replies?.length != 0) comment.replies = treeComments(comment.replies)
      return comment
    })
  }
  updateComment(newComment: lib.Comment) {
    const treeComments = (comments: lib.Comment[]) => {
      return comments.map((comment: lib.Comment) => {
        if(comment.id == newComment.id) {
          return {...comment, ...newComment}
        }
        if(comment?.replies?.length != 0) comment.replies = treeComments(comment.replies);
        return comment
      })
    }
    this.comments.all = this.comments.all.map((comment: lib.Comment) => {
      if(comment.id == newComment.id) {
        return {...comment, ...newComment}
      }
      if(comment?.replies?.length != 0) comment.replies = treeComments(comment.replies)
      return comment
    })
  }
  deleteComment(id: number) {
    const treeComments = (comments: lib.Comment[]) => {
      const newComments: lib.Comment[] = []
      comments.forEach((comment: lib.Comment) => {
        if(comment.id == id) return;
        if(comment?.replies?.length != 0) {
          comment.replies = treeComments(comment.replies);
        }
        newComments.push(comment);
      })
      return newComments
    }
    const newComments: lib.Comment[] = []
    this.comments.all.forEach((comment: lib.Comment) => {
      if(comment.id == id) return;
      if(comment?.replies?.length != 0) {
        comment.replies = treeComments(comment.replies);
      }
      newComments.push(comment);
    })
    this.comments.all = newComments
  }
}
