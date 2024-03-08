// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SvpdcSchool() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col items-center px-10 py-10 md:gap-6 sm:gap-10 gap-20 w-full\"><div class=\"flex flex-col items-center\"><h1 class=\"md:text-4xl text-lg text-sky-900 font-medium\">ST. VINCENT DE PAUL DIOCESAN COLLEGE INC.</h1></div><div class=\"flex lg:flex-row flex-col gap-10 items-center justify-center\"><div class=\"h-96 w-48 flex\"><img src=\"/assets/svpdc-mobile.svg\" alt=\"wela-school-systems\" class=\"object-fill h-96 w-44 rounded-md shadow-xl\"></div><div class=\"sm:h-full sm:w-9/12 h-40 w-screen flex flex-col items-center\"><div class=\"bg-white\"><img src=\"/assets/svpdc-school.svg\" alt=\"wela-school-systems\" class=\"object-cover h-full w-screen rounded-md shadow-xl\"></div></div></div></div><div class=\"w-auto h-full bg-white rounded-lg shadow-md px-4 py-10 flex flex-col items-center gap-10\"><div class=\"w-96 justify-center items-center gap-6 flex flex-col\"><img src=\"/assets/svpdc-logo.svg\" alt=\"wela-logo\" class=\"object-cover h-24 w-24\"> <a href=\"https://svpdc.wela.ph/\" class=\"bg-sky-400 w-32 px-4 py-2 rounded-lg text-sm text-white text-center transition ease-in-out delay-150 hover:scale-110 hover:bg-sky-700 duration-300\">View Website</a></div><p class=\"text-slate-500 md:text-2xl text-md text-left font-medium lg:w-72 w-auto\">One of the very first school assigned and task by wela to me. My very first project where I developed financial modules such as Sales, Disbursements, Inventory, Payroll, Financial Reports, Cost Centers, and Chart of Accounts. </p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
