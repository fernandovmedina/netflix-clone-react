import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"
import './index.css'
import AddProfile from './pages/addProfile.tsx'
import Browse from './pages/browse.tsx'
import CreateProfile from './pages/createProfile.tsx'
import Login from './pages/login.tsx'
import LoginHelp from './pages/LoginHelp.tsx'
import ManageProfiles from './pages/ManageProfiles.tsx'
import SelectPicture from './pages/SelectPicture.tsx'
import Signup from './pages/signup.tsx'
import CashPaymentOption from './pages/signup/cashPaymentOption.tsx'
import CreditOption from './pages/signup/creditoption.tsx'
import EditPlan from './pages/signup/editplan.tsx'
import EditPlanGift from './pages/signup/editplanGIFT.tsx'
import EditPlanOxxo from './pages/signup/editplanOXXO.tsx'
import GiftOption from './pages/signup/giftoption.tsx'
import PaymentPicker from './pages/signup/paymentPicker.tsx'
import Planform from './pages/signup/planform.tsx'
import Regform from './pages/signup/regform.tsx'
import Registration from './pages/signup/registration.tsx'
import VerifyEmail from './pages/signup/verifyemail.tsx'

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />, // verify accordion
  },
  {
    path: "/addProfile",
    element: <AddProfile />
  },
  {
    path: "/browse",
    element: <Browse />
  },
  {
    path: "/createProfile",
    element: <CreateProfile />
  },
  {
    path: "/login",
    element: <Login />
  },
  {
    path: "/loginHelp",
    element: <LoginHelp />
  },
  {
    path: "/ManageProfiles",
    element: <ManageProfiles />
  },
  {
    path: "/SelectPicture",
    element: <SelectPicture /> // needs to be checked
  },
  {
    path: "/signup",
    element: <Signup />,
  },
  {
    path: "/signup/cashPaymentOption",
    element: <CashPaymentOption />
  },
  {
    path: "/signup/creditOption",
    element: <CreditOption />
  },
  {
    path: "/signup/editplan",
    element: <EditPlan />
  },
  {
    path: "/signup/editplanGIFT",
    element: <EditPlanGift />
  },
  {
    path: "/signup/editplanOXXO",
    element: <EditPlanOxxo />
  },
  {
    path: "/signup/giftoption",
    element: <GiftOption />
  },
  {
    path: "/signup/paymentPicker",
    element: <PaymentPicker />
  },
  {
    path: "/signup/planform",
    element: <Planform />
  },
  {
    path: "/signup/regform",
    element: <Regform />
  },
  {
    path: "/signup/registration",
    element: <Registration />
  },
  {
    path: "/signup/verifyemail",
    element: <VerifyEmail />
  }
])

createRoot(document.getElementById("root")!).render(
  <RouterProvider router={router} />
)
