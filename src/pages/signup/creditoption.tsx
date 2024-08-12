import React, { useEffect } from "react";
import FooterSign from "../../components/FooterSign";
import NavbarSign from "../../components/NavbarSign";

const CreditOption = () => {
    const cardNumberURL = new URLSearchParams(window.location.search).get("cardNumber");
    const dueDateURL = new URLSearchParams(window.location.search).get("dueDate");
    const cvvURL = new URLSearchParams(window.location.search).get("cvv");
    const nameURL = new URLSearchParams(window.location.search).get("name");

    useEffect(() => {
        var inf: any = document.getElementById("inf");
        var hidden: any = document.getElementById("hidden");

        if (inf && hidden) {
            inf.addEventListener("click", () => {
                hidden.style.display = "block";
                inf.style.display = "none";
            });
        }

        var cardNumber = document.getElementById("cardNumber");
        var dueDate = document.getElementById("dueDate");
        var cvv = document.getElementById("cvv");
        var cardName = document.getElementById("cardName");

        var change: HTMLElement | any = document.getElementById("change");
        const email = new URLSearchParams(window.location.search).get("email");
        const password = new URLSearchParams(window.location.search).get("password");
        const plan = new URLSearchParams(window.location.search).get("plan");

        if (change) {
            change.addEventListener("click", () => {
                window.location.href = `/signup/editplan?email=${email}&password=${password}&plan=${plan}&cardNumber=${cardNumber.value}&dueDate=${dueDate.value}&cvv=${cvv.value}&name=${cardName.value}`;
            });
        }

        var smBTN: any = document.getElementById("smBTN");

        if (smBTN) {
            smBTN.addEventListener("click", () => {
                window.location.href = "/createProfile";
            });
        }

        const element = document.getElementById("chn");

        if (element) {
            element.addEventListener("click", () => {
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

        return () => {
            if (element) {
                element.removeEventListener("click", () => {
                    window.location.href = `/signup/paymentPicker?email=${email}&password=${password}&plan=${plan}`;
                });
            }

            if (smBTN) {
                smBTN.removeEventListener("click", () => {
                    window.location.href = "/createProfile";
                });
            }

            if (change) {
                change.removeEventListener("click", () => {
                    window.location.href = `/signup/editplan?email=${email}&password=${password}&plan=${plan}`;
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
                <div className="flex flex-col w-[35%] py-12">
                    <div className="flex mb-4 items-center">
                        <img
                            src="/behind_blue.png"
                            alt="behind_blue_image"
                            className="w-5"
                        />
                        <a
                            id="chn"
                            className="text-sm text-blue-500 hover:underline hover:underline-offset-1"
                        >Change payment method</a>
                    </div>
                    <h5 className="text-sm">
                        STEP <span className="font-bold">4</span> OF <span className="font-bold"
                        >4</span>
                    </h5>
                    <h1 className="mt-2 text-4xl font-bold">
                        Configure your credit or debit card
                    </h1>
                    <div className="flex my-3">
                        <img src="/VISA.png" alt="payment_image" />
                        <img src="/MASTERCARD.png" alt="payment_image" />
                        <img src="/AMEX.png" alt="payment_image" />
                        <img src="/CARNET.png" alt="payment_image" />
                    </div>
                    <form action="" method="post">
                        <div className="flex border-2 border-gray-500 rounded px-3 py-2">
                            <input type="text" placeholder="Card number" id="cardNumber" value={cardNumberURL} required />
                            <div className="w-full flex justify-end">
                                <img src="/card.png" alt="card_image" className="w-7" />
                            </div>
                        </div>
                        <div className="my-2 flex">
                            <input
                                type="text"
                                placeholder="Due date"
                                value={dueDateURL}
                                id="dueDate"
                                required
                                className="border-2 mr-1 border-gray-500 rounded px-3 py-2"
                            />
                            <div
                                className="flex items-center border-2 border-gray-500 rounded px-3 py-2"
                            >
                                <input
                                    type="text"
                                    placeholder="CVV"
                                    id="cvv"
                                    value={cvvURL}
                                    required
                                    className="w-[50%]"
                                />
                                <div className="w-full flex justify-end">
                                    <img
                                        src="/ayuda.png"
                                        alt="help_image"
                                        className="w-6"
                                    />
                                </div>
                            </div>
                        </div>
                        <input
                            type="text"
                            placeholder="Name in the card"
                            id="cardName"
                            required
                            value={nameURL}
                            className="w-full border-2 border-gray-500 rounded px-3 py-2"
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
                        <div className="my-2 text-xs text-gray-500">
                            <p>
                                Payments will be processed internationally. Additional
                                bank fees may apply.
                            </p>
                            <br />
                            <p>
                                By clicking the “Start Membership” button, you agree to
                                our <a
                                    href="https://help.netflix.com/legal/termsofuse"
                                    className="text-blue-500 hover:underline hover:underline-offset-1"
                                >Terms of Use</a> and <a
                                    href="https://help.netflix.com/legal/privacy"
                                    className="text-blue-500 hover:underline hover:underline-offset-1"
                                >Privacy Statement</a> and represent that you are over 18 years of age. Additionally,
                                you understand that Netflix will automatically continue your
                                membership and, until you cancel, will bill you the monthly
                                fee (currently $219 per month + applicable taxes) through
                                your chosen payment method. You can cancel your membership
                                at any time to avoid future charges. To cancel, go to Account
                                and click “Cancel Membership.”
                            </p>
                        </div>
                        <button
                            id="smBTN"
                            type="button"
                            className="bg-red-600 hover:bg-red-700 w-full rounded text-white font-bold py-3 my-3"
                        >START MEMBERSHIP</button>
                        <p className="text-xs text-gray-500">
                            This page is protected for Google reCAPTCHA to verify you
                            are not a robot. <button
                                type="button"
                                className="text-blue-600 hover:underline hover:underline-offset-1"
                                id="inf">More info</button>
                        </p>
                        <div className="hidden mt-2" id="hidden">
                            <p className="text-xs text-gray-500">
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
                    </form>
                </div>
            </section>

            <FooterSign />
        </main>
    )
}

export default CreditOption;
