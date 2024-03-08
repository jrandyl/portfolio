// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func GetInTouch() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link href=\"/styles.css\" rel=\"stylesheet\"><link rel=\"icon\" type=\"image/svg+xml\" href=\"/assets/ranaria.svg\"><script src=\"/js/htmx.min.js\"></script><title>Ranaria - JRC</title></head><body hx-boost=\"true\"><nav class=\"h-20 bg-white flex flex-row justify-between px-10 items-center w-screen shadow-md fixed\"><a href=\"/\" class=\"flex flex-row items-center\"><div class=\"md:w-30 w-30 md:items-center\"><img src=\"/assets/ranaria.svg\" class=\"w-16 h-12 cursor-pointer\" alt=\"company-logo\"></div><h1 class=\"md:flex text-sky-500 font-bold text-4xl hidden\">JRC<span class=\"text-sky-900 font-medium\">Portfolio</span></h1></a><div class=\"lg:hidden flex items-center w-12 h-12 hover:scale-110 hover:rounded-md hover:border-2 hover:border-sky-100 p-2\" hx-get=\"/side-nav\" hx-target=\"#side-nav\" hx-swap=\"outerHTML\" hx-trigger=\"click\"><img src=\"/assets/bars-solid.svg\" alt=\"bar\" class=\"w-10 h-10\"></div><div class=\"lg:flex flex-row gap-32 items-center hidden\"><ul class=\"flex justify-evenly gap-20 mt-4\"><li class=\"mx-5 cursor-pointer\"><a href=\"/\" class=\"text-sky-900 font-medium hover:border-b-2 hover:border-sky-900\">Home</a></li><li class=\"mx-5 cursor-pointer\"><a href=\"/projects\" class=\"text-sky-900 font-medium hover:border-b-2 hover:border-sky-900\">Projects</a></li><li class=\"mx-5 cursor-pointer\"><a href=\"/experience\" class=\"text-sky-900 font-medium hover:border-b-2 hover:border-sky-900\">Experience</a></li></ul><div class=\"xl:flex ml-20 hidden\"><a href=\"/get-in-touch\" class=\"w-40 h-12 flex flex-row items-center justify-center transition ease-in-out delay-150 bg-sky-400 hover:scale-110 hover:bg-sky-600 duration-300 rounded-full text-white font-medium\">Get In Touch</a></div></div></nav><div id=\"side-nav\" class=\"hidden\"></div><div id=\"root\" class=\"bg-slate-100 h-screen w-screen flex flex-col pt-10 items-center justify-center px-10 py-10\"><form action=\"/message/client\" method=\"post\" class=\"w-auto sm:px-10 sm:py-8 px-4 py-6 sm:shadow-md flex flex-col items-center gap-5 sm:w-auto sm:bg-white sm:rounded-xl\"><img src=\"/assets/ranaria.svg\" class=\"w-18 h-14\" alt=\"company-logo\"><p class=\"text-2xl font-semibold text-sky-900 mb-10\">About you</p><div><div class=\"flex md:flex-row flex-col gap-6\"><div class=\"w-full flex flex-col gap-2\"><label for=\"first_name\" class=\"text-sky-900 text-base font-medium\">First Name<span class=\"text-red-600\">*</span></label><div class=\"w-full rounded-md border-2 border-slate-300 bg-slate-200 sm:bg-slate-100 flex flex-row items-center hover:border-sky-400\"><input class=\"bg-slate-200 sm:bg-slate-100 p-2 focus:outline-none\" placeholder=\"ex. John\" name=\"first_name\" type=\"text\" required></div></div><div class=\"w-full flex flex-col gap-2\"><label for=\"middle_name\" class=\"text-sky-900 text-base font-medium\">Middle Name</label><div class=\"w-full rounded-md border-2 border-slate-300 bg-slate-200 sm:bg-slate-100 flex flex-row items-center hover:border-sky-400\"><input class=\"bg-slate-200 sm:bg-slate-100 p-2 focus:outline-none\" placeholder=\"ex. Salazar (optional)\" name=\"middle_name\" type=\"text\"></div></div></div></div><div><div class=\"flex md:flex-row flex-col gap-6\"><div class=\"w-full flex flex-col gap-2\"><label for=\"last_name\" class=\"text-sky-900 text-base font-medium\">Last Name<span class=\"text-red-600\">*</span></label><div class=\"w-full rounded-md border-2 border-slate-300 bg-slate-200 sm:bg-slate-100 flex flex-row items-center hover:border-sky-400\"><input class=\"bg-slate-200 sm:bg-slate-100 p-2 focus:outline-none\" placeholder=\"ex. Doe\" name=\"last_name\" type=\"text\" required></div></div><div class=\"w-full flex flex-col gap-2\"><label for=\"email\" class=\"text-sky-900 text-base font-medium\">Email<span class=\"text-red-600\">*</span></label><div class=\"w-full rounded-md border-2 border-slate-300 bg-slate-200 sm:bg-slate-100 flex flex-row items-center hover:border-sky-400\"><input class=\"bg-slate-200 sm:bg-slate-100 p-2 focus:outline-none\" placeholder=\"ex. user@example.com\" name=\"email\" type=\"email\" required></div></div></div></div><div class=\"w-full flex flex-col gap-2\"><label for=\"message\" class=\"text-sky-900 text-base font-medium\">General message</label> <textarea name=\"message\" placeholder=\"Type your message here...\" class=\"w-full bg-slate-200 sm:bg-slate-100 py-4 px-6 rounded-md border-2 border-slate-300 bg-slate-200 sm:bg-slate-100 hover:border-sky-400\"></textarea></div><button type=\"submit\" class=\"h-12 px-10 items-center text-white font-medium bg-sky-400 rounded-md hover:bg-sky-400/50 hover:text-white\">Submit</button></form></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
