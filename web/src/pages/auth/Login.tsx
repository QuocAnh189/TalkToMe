//hooks
import { useWindowSize } from 'react-use'

//assets
import auth_img from '@assets/images/background_auth.jpg'
import logo_img from '@assets/images/logo.png'
import FormLogin from '@components/auth/FormSignIn'

const Login = () => {
  const { width } = useWindowSize()

  return (
    <div className="flex-1 grid grid-cols-1 xl:grid-cols-2 4xl:grid-cols-[minmax(0,_1030px)_minmax(0,_1fr)]">
      {width >= 1280 && (
        <div className="flex flex-col justify-center items-center lg:p-[60px] bg-[#e3f2fd]">
          <div className="flex items-center">
            <a className="logo" href="/">
              <img loading="lazy" src={logo_img} alt="EventHub" className="w-[200px] object-cover" />
            </a>
            <p className="text-center text-header tracking-[0.2px] font-semibold text-xl leading-6 max-w-[540px] my-7 mx-auto">
              Welcome to Go Chat
            </p>
          </div>
          <img loading="lazy" className="" src={auth_img} alt="media" />
        </div>
      )}
      <div className="relative w-full h-screen flex justify-center items-center">
        <main className="mx-auto mt-auto flex min-h-screen w-full max-w-full flex-col overflow-hidden bg-[#6cb2eb]">
          <div className="absolute left-[50%] top-[50%] w-[600px] translate-x-[-50%] translate-y-[-50%] overflow-hidden rounded-[50px] bg-[#e3f2fd] px-[100px] py-[60px] mdl:min-h-[600px]">
            <div className="mb-[30px] flex flex-row items-center justify-center gap-x-4">
              <h1 className="text-4xl font-semibold">Login</h1>
            </div>
            <FormLogin />
          </div>
        </main>
      </div>
    </div>
  )
}

export default Login
