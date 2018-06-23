import auth0 from 'auth0-js'

const clientID = 'BalIDVTaK374gSwltgrsIQFWakihPSI9'
const domain = 'zenika.eu.auth0.com'

export const webAuth = new auth0.WebAuth({
  domain,
  clientID,
  redirectUri: window.location.href,
  audience: 'https://grao.zenika.com/api/v1',
  responseType: 'token id_token',
  scope: 'openid',
  leeway: 60
})

export const bootstrapAuth0 = () => {
  webAuth.parseHash((err, authResult) => {
    if (authResult && authResult.accessToken && authResult.idToken) {
      window.location.hash = ''
      setSession(authResult)
    } else if (err) {
      console.log(err)
    }
  })
}

export const setSession = (authResult) => {
  var expiresAt = JSON.stringify(
    authResult.expiresIn * 1000 + new Date().getTime()
  )
  localStorage.setItem('access_token', authResult.accessToken)
  localStorage.setItem('id_token', authResult.idToken)
  localStorage.setItem('expires_at', expiresAt)
}

export const isConnected = () => {
  const expiresAt = JSON.parse(localStorage.getItem('expires_at'))
  return new Date().getTime() < expiresAt
}

export const getAuthHeader = () => {
  return {
    'Authorization': 'Bearer ' + localStorage.getItem('access_token')
  }
}

export const login = () => {
  webAuth.authorize()
}

export const logout = () => {
  localStorage.removeItem('id_token')
  localStorage.removeItem('expires_at')
  localStorage.removeItem('profile')
}
