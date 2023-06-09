
import { createStore,createLogger } from 'vuex'
 
export default createStore({
  state: {
    movieID:'3',
    userList:[]
  },
  // getters: {
  //   getUserList(state){
  //     return state.userList
  //   }
  // },
  mutations: {
    getMovieId(state,movieID){
      state.movieID = movieID
    },
    setUserList(state,userList){
      state.userList = userList
    }
  },
  actions: {
    
  },
  modules: {
  },
  getters:{
    getUserList(state) {
      console.log(state.userList);
      return state.userList
    }
  }
})