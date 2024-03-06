// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package fragments

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Project() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"h-20 bg-white flex flex-row justify-between px-10 items-center w-screen shadow-md fixed\"><a href=\"/\" class=\"flex flex-row items-center\"><div class=\"md:w-30 w-30 md:items-center\"><img src=\"/assets/ranaria.svg\" class=\"w-16 h-12 cursor-pointer\" alt=\"company-logo\"></div><h1 class=\"text-sky-500 font-bold text-4xl\">JRC <span class=\"text-sky-900 font-medium\">Porfolio</span></h1></a><div class=\"lg:hidden h-6 w-10 cursor-pointer\"><img src=\"/assets/bar.svg\" alt=\"bar\"></div><div class=\"lg:flex flex-row gap-32 items-center hidden\"><ul class=\"flex justify-evenly gap-20 mt-4\"><li class=\"mx-5 cursor-pointer\"><a href=\"/\" class=\"text-sky-900 font-medium hover:border-b-2 hover:border-sky-900\">Home</a></li><li class=\"mx-5 cursor-pointer\"><a href=\"/projects\" class=\"text-sky-900 font-medium hover:border-b-2 hover:border-sky-900\">Projects</a></li><li class=\"mx-5 cursor-pointer\"><a href=\"/experience\" class=\"text-sky-900 font-medium hover:border-b-2 hover:border-sky-900\">Experience</a></li></ul><div class=\"xl:flex ml-20 hidden\"><a href=\"#/get-in-touch\" class=\"w-40 h-12 flex flex-row items-center justify-center transition ease-in-out delay-150 bg-sky-400 hover:scale-110 hover:bg-sky-600 duration-300 rounded-full text-white font-medium\">Get In Touch</a></div></div></nav><div class=\"bg-slate-100 lg:h-full h-full lg:w-screen w-screen flex flex-col pt-10 gap-10 items-center px-20 py-20\"><h1 class=\"mt-20 text-sky-900 font-bold text-6xl\">Featured Works</h1><div class=\"w-full h-full flex flex-row gap-6\"><div class=\"flex flex-col items-center px-10 py-10 gap-6 w-full\"><div class=\"flex flex-row gap-10 items-center justify-center\"><div class=\"h-96 w-48 flex\"><img src=\"/assets/wela-mobile.svg\" alt=\"wela-school-systems\" class=\"object-fill h-96 w-44 rounded-md shadow-xl\"></div><div class=\"h-full w-9/12 flex flex-col items-center\"><div class=\"bg-white\"><img src=\"/assets/wela.svg\" alt=\"wela-school-systems\" class=\"object-cover h-full w-full rounded-md shadow-xl\"></div></div></div><div class=\"flex flex-col items-center\"><h1 class=\"text-4xl text-sky-900 font-medium\">WELA SCHOOL SYSTEM</h1></div></div><div class=\"w-auto h-full bg-white rounded-lg shadow-md px-4 py-10 flex flex-col items-center gap-10\"><div class=\"w-96 justify-center items-center gap-6 flex flex-col\"><img src=\"/assets/wela-logo.svg\" alt=\"wela-logo\" class=\"object-cover h-12 w-44\"> <a href=\"https://wela.online/\" class=\"bg-sky-400 w-32 px-4 py-2 rounded-lg text-sm text-white text-center transition ease-in-out delay-150 hover:scale-110 hover:bg-sky-700 duration-300\">View Website</a></div><p class=\"text-slate-500 text-2xl text-left font-medium w-72\">For the past two years, my primary focus has revolved around crafting sophisticated systems that streamline various aspects of school management, with a particular emphasis on Billing, Finance, and Accounting functionalities.</p></div></div><div class=\"w-full h-full flex flex-row gap-6 mt-40\"><div class=\"flex flex-col items-center px-10 py-10 gap-6 w-full\"><div class=\"flex flex-row gap-10 items-center justify-center\"><div class=\"h-96 w-48 flex\"><img src=\"/assets/claret-school-mobile.svg\" alt=\"wela-school-systems\" class=\"object-fill h-96 w-44 rounded-md shadow-xl\"></div><div class=\"h-full w-9/12 flex flex-col items-center\"><div class=\"bg-white\"><img src=\"/assets/claret-school.svg\" alt=\"wela-school-systems\" class=\"object-cover h-full w-full rounded-md shadow-xl\"></div></div></div><div class=\"flex flex-col items-center\"><h1 class=\"text-4xl text-sky-900 font-medium\">CLARET SCHOOL OF ZAMBOANGA</h1></div></div><div class=\"w-auto h-full bg-white rounded-lg shadow-md px-4 py-10 flex flex-col items-center gap-10\"><div class=\"w-96 justify-center items-center gap-6 flex flex-col\"><img src=\"/assets/claret-school-logo.svg\" alt=\"wela-logo\" class=\"object-cover h-24 w-24\"> <a href=\"https://claretzambo.wela.ph/\" class=\"bg-sky-400 w-32 px-4 py-2 rounded-lg text-sm text-white text-center transition ease-in-out delay-150 hover:scale-110 hover:bg-sky-700 duration-300\">View Website</a></div><p class=\"text-slate-500 text-2xl text-left font-medium w-72\">I've played a pivotal role in the development and refinement of essential financial modules such as Sales, Disbursements, Inventory, Payroll, Financial Reports, Cost Centers, and Chart of Accounts within Claret School.</p></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}