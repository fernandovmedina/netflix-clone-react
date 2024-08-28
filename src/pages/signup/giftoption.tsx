import { useEffect } from "react";
import FooterSign from "../../components/FooterSign";
import NavbarSign from "../../components/NavbarSign";

const GiftOption = () => {
    const email = new URLSearchParams(window.location.search).get("email");
    const password = new URLSearchParams(window.location.search).get("password");
    const plan = new URLSearchParams(window.location.search).get("plan");
    const giftCodeURL = new URLSearchParams(window.location.search).get("giftCode")

    useEffect(() => {
        const element = document.getElementById("chn")
        const arrow: any = document.getElementById("arrow_back");
        const giftCode = document.getElementById("giftCode");

        if (element) {
            element.addEventListener("click", () => {
                window.location.href = `/signup/paymentPicker?email=${email}&password=${password}&plan=${plan}`;
            });
        }

        if (arrow) {
            arrow.addEventListener('click', () => {
                window.location.href = `/signup/paymentPicker?email=${email}&password=${password}&plan=${plan}`;
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

        var inf: any = document.getElementById("inf");
        var hidden: any = document.getElementById("hidden");

        if (inf && hidden) {
            inf.addEventListener("click", () => {
                hidden.style.display = "block";
                inf.style.display = "none";
            });
        }

        var change: HTMLElement | any = document.getElementById("change");

        if (change) {
            change.addEventListener("click", () => {
                window.location.href = `/signup/editplanGIFT?email=${email}&password=${password}&plan=${plan}&giftCode=${giftCode.value}`;
            });
        }

        return () => {
            if (element) {
                element.removeEventListener("click", () => {
                    window.location.href = `/signup/paymentPicker?email=${email}&password=${password}&plan=${plan}`;
                });
            }

            if (arrow) {
                arrow.removeEventListener('click', () => {
                    window.location.href = `/signup/paymentPicker?email=${email}&password=${password}&plan=${plan}`
                });
            }

            if (change) {
                change.removeEventListener("click", () => {
                    window.location.href = `/signup/editplanGIFT?email=${email}&password=${password}&plan=${plan}`;
                });
            }

            if (inf && hidden) {
                inf.removeEventListener("click", () => {
                    hidden.style.display = "block";
                    inf.style.display = "none";
                });
            }
        }
    }, []);

    return (
        <main>
            <NavbarSign />

            <section className="w-full flex items-center justify-center">
                <div className="py-12 w-[35%]">
                    <div className="flex mb-4 align-center w-full">
                        <div className="w-1/2 flex items-center">
                            <img
                                src="/behind_blue.png"
                                alt="behind_blue_image"
                                className="w-5"
                                id="arrow_back"
                            />
                            <a
                                id="chn"
                                className="text-blue-500 hover:underline hover:underline-offset-1"
                            >Change payment method</a>
                        </div>
                        <div className="w-1/2 flex justify-end">
                            <a className="text-blue-500 hover:underline hover:underline-offset-1">Generate new gift code</a>
                        </div>
                    </div>
                    <h5 className="text-sm">
                        STEP <span className="font-bold">4</span> OF <span className="font-bold"
                        >4</span>
                    </h5>
                    <h1 className="mt-2 text-4xl font-bold">Enter your gift card</h1>
                    <input
                        type="text"
                        placeholder="PIN or gift code"
                        id="giftCode"
                        value={giftCodeURL}
                        required
                        className="border-2 border-gray-400 px-3 py-2 w-full my-2"
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
                            >CHANGE</a>
                        </div>
                    </div>
                    <button
                        type="button"
                        className="w-full text-white bg-red-600 hover:bg-red-700 rounded py-3 my-2"
                    >REDEEM GIFT CODE</button>
                    <div>
                        <p className="text-sm text-gray-500">
                            This page is protected for Google reCAPTCHA to verify you
                            are not a robot. <button
                                type="button"
                                className="text-blue-600 hover:underline hover:underline-offset-1"
                                id="inf">More info</button>
                        </p>
                        <div className="hidden mt-2" id="hidden">
                            <p className="text-sm text-gray-500">
                                The information collected by Google reCAPTCHA is subject
                                to Google's <a
                                    className="text-blue-600 hover:underline hover:underline-offset-1"
                                    href="https://policies.google.com/privacy"
                                >Privacy Policy</a> and <a
                                    href="https://policies.google.com/terms"
                                    className="text-blue-600 hover:underline hover:underline-offset-1"
                                >Terms of Service</a>, and is used to provide, maintain and improve the
                                reCAPTCHA service, as well as for general security
                                purposes (Google does not use it to personalize
                                advertising).
                            </p>
                        </div>
                    </div>
                </div>
            </section>

            <FooterSign />
        </main>
    )
}

export default GiftOption;
