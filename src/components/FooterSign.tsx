const FooterSign = () => {
  return (
    <footer
      className="w-full bg-gray-100 text-gray-600 px-16 py-12 border border-t-2 border-gray-900"
    >
      <h4 className="font-normal mb-4">
        Questions? Call to <a href="tel:800 953 1430">800 953 1430</a>
      </h4>
      <div className="w-full flex text-sm underline underline-offset-1 pb-4">
        <div className="w-1/4 flex flex-col">
          <a href="" className="hover:text-black">Frecuent questions</a>
          <a href="" className="hover:text-black">Privacy</a>
        </div>
        <div className="w-1/4 flex flex-col">
          <a href="" className="hover:text-black">Help center</a>
          <a href="" className="hover:text-black">Cookies preferences</a>
        </div>
        <div className="w-1/4 flex flex-col">
          <a href="" className="hover:text-black">Netflix store</a>
          <a href="" className="hover:text-black">Corporative information</a>
        </div>
        <div className="w-1/4 flex flex-col">
          <a href="" className="hover:text-black">Terms of use</a>
        </div>
      </div>
      <button
        className="flex items-center mx-2 bg-black mb-4 text-white border border-gray-500 px-2 py-1 rounded"
      >
        <img
          src="/icons8-translate-google-64.png"
          alt="google_translate_img"
          className="w-6"
        />
        <select name="languages" id="languages" className="bg-black text-white">
          <option value="english">English</option>
          <option value="spanish">Spanish</option>
        </select>
      </button>
      <p>Netflix</p>
    </footer>
  )
}

export default FooterSign;
