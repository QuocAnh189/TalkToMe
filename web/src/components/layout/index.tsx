//components
import { Outlet } from 'react-router-dom'
import { Toaster } from 'react-hot-toast'

const Layout = () => {
  return (
    <div className="">
      <Outlet />
      <Toaster />
    </div>
  )
}

export default Layout
