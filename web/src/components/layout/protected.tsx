//hooks
import AppSocketIOProvider from 'context/socket_io'
import { ReactNode, useLayoutEffect } from 'react'
import { useNavigate } from 'react-router-dom'

const ProtectedLayout = ({ children }: { children: ReactNode }) => {
  const navigate = useNavigate()
  const user = localStorage.getItem('user')

  useLayoutEffect(() => {
    if (user === null) {
      navigate('/login')
    }
  }, [user])

  return <AppSocketIOProvider>{children}</AppSocketIOProvider>
}

export default ProtectedLayout
