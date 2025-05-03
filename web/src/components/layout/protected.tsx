//hooks
import { ReactNode } from 'react'
// import { useNavigate } from 'react-router-dom'

const ProtectedLayout = ({ children }: { children: ReactNode }) => {
  // const navigate = useNavigate()
  // const user = localStorage.getItem('user')

  // useLayoutEffect(() => {
  //   if (user === null) {
  //     navigate('/login')
  //   }
  // }, [user])

  return <div>{children}</div>
}

export default ProtectedLayout
