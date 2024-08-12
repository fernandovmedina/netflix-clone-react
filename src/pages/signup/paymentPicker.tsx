import NavbarSign from "../../components/NavbarSign";
import FooterSign from "../../components/FooterSign";
import { useEffect } from "react";

const PaymentPicker = () => {
    const email = new URLSearchParams(window.location.search).get("email");
    const password = new URLSearchParams(window.location.search).get("password");
    const plan = new URLSearchParams(window.location.search).get("plan");

    useEffect(() => {
        var payment_1: any = document.getElementById("payment_1");
        var payment_2: any = document.getElementById("payment_2");
        var payment_3: any = document.getElementById("payment_3");

        if (payment_1 && payment_2 && payment_3) {
            payment_1.addEventListener("click", () => {
                window.location.href = `/signup/creditoption?email=${email}&password=${password}&plan=${plan}`;
            });

            payment_2.addEventListener("click", () => {
                window.location.href = `/signup/cashPaymentOption?email=${email}&password=${password}&plan=${plan}`;
            });

            payment_3.addEventListener("click", () => {
                window.location.href = `/signup/giftoption?email=${email}&password=${password}&plan=${plan}`;
            });
        }

        return () => {
            if (payment_1 && payment_2 && payment_3) {
                payment_1.removeEventListener("click", () => {
                    window.location.href = `/signup/creditoption?email=${email}&password=${password}&plan=${plan}`;
                });

                payment_2.removeEventListener("click", () => {
                    window.location.href = `/signup/cashPaymentOption?email=${email}&password=${password}&plan=${plan}`;
                });
    
                payment_3.removeEventListener("click", () => {
                    window.location.href = `/signup/giftoption?email=${email}&password=${password}&plan=${plan}`;
                });
            }
        }
    }, []);

    return (
        <main>
            <NavbarSign />

            <section className="py-10">
                <div className="flex flex-col items-center justify-center">
                    <div className="w-[35%] text-center">
                        <img src="/candado.png" className="mx-auto" alt="candado_image" />
                        <h5 className="text-sm mt-8">
                            STEP <span className="font-bold">4</span> OF <span
                                className="font-bold">4</span>
                        </h5>
                        <h1 className="font-bold text-4xl">Choose how you wanna pay</h1>
                        <h3 className="mt-3 text-lg">
                            Your payment is encrypted and you can change it whenever you
                            want
                        </h3>
                        <h3 className="mt-2 text-lg font-semibold">
                            Safe transactions and reliable. Cancel easy online.
                        </h3>
                        <div className="w-full flex justify-end items-center mt-3">
                            <h6 className="text-sm">End-to-end encryption</h6>
                            <img
                                className="w-6"
                                src="/candado_amarillo.png"
                                alt="candado_amarillo_image"
                            />
                        </div>
                        <div
                            id="payment_1"
                            className="flex items-center w-full border-2 border-gray-400 rounded px-3 py-3 cursor-pointer mb-2"
                        >
                            <div className="w-5/6 flex">
                                <h4 className="mr-3">Credit or debit card</h4>
                                <img className="mx-1" src="/VISA.png" alt="payment_image" />
                                <img
                                    className="mx-1"
                                    src="/MASTERCARD.png"
                                    alt="payment_image"
                                />
                                <img className="mx-1" src="/AMEX.png" alt="payment_image" />
                                <img
                                    className="mx-1"
                                    src="/CARNET.png"
                                    alt="payment_image"
                                />
                            </div>
                            <div className="w-1/6 flex justify-end">
                                <img
                                    src="/adelante_negro.png"
                                    alt="adelante_negro_image"
                                    className="w-5"
                                />
                            </div>
                        </div>
                        <div
                            id="payment_2"
                            className="flex items-center w-full border-2 border-gray-400 rounded px-3 py-3 cursor-pointer mb-2"
                        >
                            <div className="w-5/6 flex items-center">
                                <h4 className="mr-2">OXXO PAY (MEXICO)</h4>
                                <img src="/OXXO.png" alt="payment_image" />
                            </div>
                            <div className="w-1/6 flex justify-end">
                                <img
                                    src="/adelante_negro.png"
                                    alt="adelante_negro_image"
                                    className="w-5"
                                />
                            </div>
                        </div>
                        <div
                            id="payment_3"
                            className="flex items-center w-full border-2 border-gray-400 rounded px-3 py-3 cursor-pointer"
                        >
                            <div className="w-5/6 flex items-center">
                                <h4 className="mr-2">Gift card (OBVIOUSLY DOESN'T WORK)</h4>
                                <img src="/GIFT_CODE.png" alt="payment_image" />
                            </div>
                            <div className="w-1/6 flex justify-end">
                                <img
                                    src="/adelante_negro.png"
                                    alt="adelante_negro_image"
                                    className="w-5"
                                />
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <FooterSign />
        </main>
    )
}

export default PaymentPicker;
