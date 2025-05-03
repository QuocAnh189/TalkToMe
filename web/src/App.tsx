import { Suspense, lazy } from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'

//components
import Layout from '@components/layout'
import Loader from '@components/common/Loader'

///screens
const LoginPage = lazy(() => import('@pages/auth/Login'))
const RegisterPage = lazy(() => import('@pages/auth/Register'))
const HomePage = lazy(() => import('@pages/home/Home'))
const ProfilePage = lazy(() => import('@pages/profile/Profile'))
const UserPage = lazy(() => import('@pages/user/User'))
const NotFound = lazy(() => import('@pages/errors/NotFound'))

function App() {
  return (
    <BrowserRouter>
      <Suspense fallback={<Loader />}>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route path="" element={<HomePage />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/profile" element={<ProfilePage />} />
            <Route path="/user/:id" element={<UserPage />} />
          </Route>
          <Route path="/*" element={<NotFound />} />
        </Routes>
      </Suspense>
    </BrowserRouter>
  )
}

export default App
