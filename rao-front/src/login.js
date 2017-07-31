import Auth0Lock from 'auth0-lock'
import Vue from 'vue'

const clientID = 'BalIDVTaK374gSwltgrsIQFWakihPSI9'
const clientDomain = 'zenika.eu.auth0.com'

export const lock = new Auth0Lock(clientID, clientDomain)

export const isConnected = () => {
  return !!localStorage.getItem('id_token')
}

export const getAuthHeader = () => {
  return {
    'Authorization': 'Bearer ' + localStorage.getItem('access_token')
  }
}

export const login = () => {
  // Show the Lock Widget and save the user's JWT on a successful login
  lock.show((err, profile, idToken) => {
    if (err) {
      console.error(err)
      login()
      return
    }
    localStorage.setItem('profile', JSON.stringify(profile))
    localStorage.setItem('id_token', idToken)
    Vue.http.headers.common['Authorization'] = getAuthHeader()
  })
}

export const logout = () => {
  localStorage.removeItem('profile')
  localStorage.removeItem('id_token')
}
