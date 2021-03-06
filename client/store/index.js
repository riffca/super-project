import Vue from 'vue'
import Vuex from 'vuex'
import * as userService from 'root/services/user'
import * as leadService from 'root/services/lead'

Vue.use(Vuex)

const state = {
  user: {
    all:{},
    profile:{},
    id:'',
    authId:'',
  },
  leads:{
    all:[],
    id:''
  },
  count: 0,
  FUNC_ARGUMENT:{}
}


const getters = {
  userOne: state=>state.user.all[state.user.id],
  userAll: state =>state.user.all,
  userAuth: state=>state.user.all[state.user.authId]
}

const mutations = {
  USER_SET_AUTH(state) {
    let user=state.FUNC_ARGUMENT
    state.user.authId = user.user_name
    state.user.all[user.user_name]=user
  },
  USER_VIEW_ONE(state) {
    let user=state.FUNC_ARGUMENT
    state.user.all[user.user_name]=user
    state.user.id = user.user_name
  },
  USER_VIEW_ALL(state) {
    let users = state.FUNC_ARGUMENT
    users.forEach((user,index)=>{
      state.user.all[user.user_name]=user
    })
  },
  LEAD_PUSH_ONE(state){
    let lead = state.FUNC_ARGUMENT
    state.leads.all[lead.id]=lead
  },
  LEAD_SET_USER_LEADS(state){
    let leads = state.FUNC_ARGUMENT
    console.log(leads)
    state.leads.all = leads
  },
  INCREMENT (state) {
    state.count++
  },
  DECREMENT (state) {
    state.count--
  }
}

const actions = {
  authUser({state,commit},{user}){
    state.FUNC_ARGUMENT=user
    commit("USER_SET_AUTH")
  },
  incrementAsync ({ state, commit }) {
    setTimeout(() => {
      commit('INCREMENT')
    }, 200)
  },
  getUsers({ state, commit }, { routeParam="" }){
    userService
      .GetOne({user_name: routeParam})
      .then(data=>{
        state.FUNC_ARGUMENT = data.Value
        commit("USER_VIEW_ONE")
      })
    userService
      .GetAll()
      .then(data=>{
        state.FUNC_ARGUMENT = data.Value
        commit("USER_VIEW_ALL")
      })
  },
  addLead({state,commit},{adress}){
    leadService
      .createLead({
        adress_id:adress.id+"",
        creator_id: state.user.authId+"",
        status_code: "10"
      })
      .then(data=>{
        state.FUNC_ARGUMENT = data.Value
         commit("LEAD_PUSH_ONE")
      })
  },
  leadSet({state,commit}){
    leadService
      .getUserLeads({creator_id: "1"})
      .then(data=>{
        state.FUNC_ARGUMENT=data.Value
        commit("LEAD_SET_USER_LEADS")
      })
  }
}

const store = new Vuex.Store({
  state,
  mutations,
  actions,
  getters
})

export default store
