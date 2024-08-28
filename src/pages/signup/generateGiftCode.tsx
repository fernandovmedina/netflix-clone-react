import { useEffect } from 'react'
import FooterSign from '../../components/FooterSign.tsx'
import NavbarSign from '../../components/NavbarSign.tsx'

const GenerateGiftCode = () => {
    
    useEffect(() => {
        const giftCode: HTMLElement | any = document.getElementById("gift_code");

        // call to api to get a new gift code
    }, [])

    return (
        <main>
            <div>
                <NavbarSign />
            </div>
            <section className='w-full flex items-center justify-center my-20'>
                <div className='flex flex-col items-center'>
                    <h1 className='text-3xl font-bold'>Random Code Generator</h1>
                    <p className='mt-5 font-medium text-lg'>The generated codes will be stored in our database and you are going to be able to use it as a payment method</p>
                    <h1 id='gift_code'></h1>
                    <button className='bg-red-600 rounded mt-10 px-5 py-3 text-white font-bold text-lg hover:bg-red-700 hover:text-black'>Generate</button>
                </div>
            </section>
            <div>
                <FooterSign />
            </div>
        </main>
    )
}

export default GenerateGiftCode;
