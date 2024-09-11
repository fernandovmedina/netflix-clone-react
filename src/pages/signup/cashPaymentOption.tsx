import { useEffect } from "react";
import FooterSign from "../../components/FooterSign";
import NavbarSign from "../../components/NavbarSign";

const CashPaymentOption = () => {
  const email = new URLSearchParams(window.location.search).get("email");
  const password = new URLSearchParams(window.location.search).get("password");
  const plan = new URLSearchParams(window.location.search).get("plan");
  const phoneNumberURL = new URLSearchParams(window.location.search).get(
    "phoneNumber"
  );
  const clientNameURL = new URLSearchParams(window.location.search).get(
    "clientName"
  );

  useEffect(() => {
    document.getElementById("chn")!.addEventListener("click", () => {
      window.location.href = `/signup/paymentPicker?email=${email}&password=${password}&plan=${plan}`;
    });

    var clientName: string | any = document.getElementById("clientName");
    var phoneNumber: string | any = document.getElementById("phoneNumber");

    var chk: any = document.getElementById("chk");
    var check: any = document.getElementById("checked");

    if (chk && check) {
      chk.addEventListener("click", () => {
        check.checked = !check.checked;
      });
    }

    var plans = {
      premium: {
        name: "Premium",
        price: "299",
      },
      standard: {
        name: "Standard",
        price: "219",
      },
      ads: {
        name: "Standard with ads",
        price: "99",
      },
    };

    var planName: any = document.getElementById("planName");
    var planPrice: any = document.getElementById("planPrice");

    if (planName && planPrice) {
      if (plan === "premium") {
        planName.innerHTML = plans.premium.name;
        planPrice.innerHTML = plans.premium.price;
      } else if (plan === "standard") {
        planName.innerHTML = plans.standard.name;
        planPrice.innerHTML = plans.standard.price;
      } else {
        planName.innerHTML = plans.ads.name;
        planPrice.innerHTML = plans.ads.price;
      }
    }

    var change: any = document.getElementById("change");

    if (change) {
      change.addEventListener("click", () => {
        window.location.href = `/signup/editplanOXXO?email=${email}&password=${password}&plan=${plan}&clientName=${clientName.value}&phoneNumber=${phoneNumber.value}`;
      });
    }

    return () => {
      if (change) {
        change.removeEventListener("click", () => {
          window.location.href = `/signup/editplanOXXO?email=${email}&password=${password}&plan=${plan}`;
        });
      }

      if (chk && check) {
        chk.removeEventListener("click", () => {
          check.checked = !check.checked;
        });
      }
    };
  }, []);

  return (
    <main>
      <NavbarSign />

      <section className="w-full flex items-center justify-center">
        <div className="flex flex-col w-[30%] py-12">
          <div className="flex mb-4 items-center w-full">
            <div className="w-1/2 flex items-center">
              <img
                src="/behind_blue.png"
                alt="behind_blue_image"
                className="w-5"
              />
              <a
                id="chn"
                className="text-sm text-blue-500 hover:underline hover:underline-offset-1"
              >
                Change payment method
              </a>
            </div>
            <div className="w-1/2 flex justify-end">
                <a href="/signup/oxxoCodeGenerator" className="items-center text-sm text-blue-500 hover:underline hover:underline-offset-1">Generate Code</a>
            </div>
          </div>
          <h1 className="font-bold text-4xl mb-6">
            Get your OXXO PAY referral code
          </h1>
          <h4 className="text-lg mb-6">
            We'll send you the code to this phone number.
          </h4>
          <h4 className="text-base">
            Your number will also be used in case you forget your password and
            for important account messages. SMS rates may apply.
          </h4>
          <form action="" method="post" className="my-6">
            <div className="flex items-center border-2 border-gray-600 p-2">
              <img src="/mexico.png" alt="mexico_flag_image" className="w-8" />
              <h5 className="mx-2">+52</h5>
              <input
                type="tel"
                placeholder="Phone number"
                id="phoneNumber"
                value={phoneNumberURL}
                className="w-full"
                required
              />
            </div>
            <input
              type="text"
              placeholder="Name"
              id="clientName"
              value={clientNameURL}
              required
              className="border-2 border-gray-600 my-2 w-full px-2 py-3"
            />
            <input
                type="text"
                placeholder="Code"
                id="oxxoCode"
                required
                className="border-2 border-gray-600 w-full px-2 py-3"
            />
            <div className="bg-gray-200 p-3 flex my-2 rounded">
              <div className="w-1/2">
                <h1 className="font-bold text-base">
                  $<span id="planPrice"></span> monthly
                </h1>
                <h2 className="font-light" id="planName"></h2>
              </div>
              <div className="w-1/2 flex justify-end items-center">
                <a
                  id="change"
                  className="text-blue-500 font-bold text-sm hover:underline hover:underline-offset-1"
                >
                  CHANGE
                </a>
              </div>
            </div>
            <div className="text-gray-400 text-xs">
              <p>
                By clicking the “Start Membership” button, you agree to our{" "}
                <a
                  className="text-blue-600 hover:underline hover:underline-offset-1"
                  href="https://help.netflix.com/legal/termsofuse"
                >
                  Terms of Use
                </a>{" "}
                and{" "}
                <a
                  className="text-blue-600 hover:underline hover:underline-offset-1"
                  href="https://help.netflix.com/legal/privacy"
                >
                  Privacy Statement
                </a>{" "}
                and represent that you are over 18 years of age. Additionally,
                you understand that Netflix will automatically continue your
                membership and, until you cancel, will bill you the monthly fee
                (currently $99 per month + applicable taxes) through your chosen
                payment method. You can cancel your membership at any time to
                avoid future charges. To cancel, go to Account and click “Cancel
                Membership.”
              </p>
            </div>
            <div className="flex items-center my-2">
              <input type="checkbox" required id="checked" />
              <h4 className="mx-2 text-base" id="chk">
                I ACCEPT
              </h4>
            </div>
            <button
              type="button"
              className="w-full mt-8 bg-red-600 text-white py-3 rounded hover:bg-red-700"
            >
              START MEMBERSHIP
            </button>
          </form>
        </div>
      </section>

      <FooterSign />
    </main>
  );
};

export default CashPaymentOption;
