import * as Type from './type'
export class Comments {
    comments = {
    all: [] as Type.Comment[],
  }
  insertComments(comments: Type.Comment[]) {
    this.comments.all = comments;
  }
  addComment(newComment: Type.Comment) {
    if(newComment.parent_id == 0) {
      this.comments.all.push(newComment);
      return
    }
    const treeComments = (comments: Type.Comment[]) => {
      return comments.map((comment: Type.Comment) => {
        if(comment.id == newComment.parent_id) {
          comment.replies.push(newComment);
          return comment
        }
        if(comment?.replies?.length != 0) comment.replies = treeComments(comment.replies);
        return comment
      })
    }
    this.comments.all = this.comments.all.map((comment: Type.Comment) => {
      if(comment.id == newComment.parent_id) {
        comment.replies.push(newComment);
        return comment
      }
      if(comment?.replies?.length != 0) comment.replies = treeComments(comment.replies)
      return comment
    })
  }
  updateComment(newComment: Type.Comment) {
    const treeComments = (comments: Type.Comment[]) => {
      return comments.map((comment: Type.Comment) => {
        if(comment.id == newComment.id) {
          return {...comment, ...newComment}
        }
        if(comment?.replies?.length != 0) comment.replies = treeComments(comment.replies);
        return comment
      })
    }
    this.comments.all = this.comments.all.map((comment: Type.Comment) => {
      if(comment.id == newComment.id) {
        return {...comment, ...newComment}
      }
      if(comment?.replies?.length != 0) comment.replies = treeComments(comment.replies)
      return comment
    })
  }
  deleteComment(id: number) {
    const treeComments = (comments: Type.Comment[]) => {
      const newComments: Type.Comment[] = []
      comments.forEach((comment: Type.Comment) => {
        if(comment.id == id) return;
        if(comment?.replies?.length != 0) {
          comment.replies = treeComments(comment.replies);
        }
        newComments.push(comment);
      })
      return newComments
    }
    const newComments: Type.Comment[] = []
    this.comments.all.forEach((comment: Type.Comment) => {
      if(comment.id == id) return;
      if(comment?.replies?.length != 0) {
        comment.replies = treeComments(comment.replies);
      }
      newComments.push(comment);
    })
    this.comments.all = newComments
  }
}
