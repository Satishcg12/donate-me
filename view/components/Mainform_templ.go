// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Mainform() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"flex flex-col gap-5\" hx-post=\"/api/coffee\" hx-swap=\"outerHTML\" x-data=\"{ total:100 }\"><div class=\"flex justify-center  items-center text-gray-500 p-2 gap-5  bg-blue-100 border border-gray-300 shadow shadow-blue-100  rounded-2xl\" x-data><div class=\"flex items-center gap-2\"><img src=\"/static/images/coffee.png\" alt=\"coffee\" class=\"h-16 w-16 rounded-full\" title=\"Buy me a coffee to support my work worth 100 NRS\"> X</div><div><input type=\"radio\" value=\"1\" name=\"amount\" id=\"one\" class=\"hidden peer\" checked @change=\"$refs.custominput.value =0; total=100 \"> <label for=\"one\" class=\"grid place-content-center rounded-full border border-blue-400 bg-blue-50 peer-checked:bg-blue-500 peer-checked:text-white h-10 w-10\">1</label></div><div><input type=\"radio\" value=\"2\" name=\"amount\" id=\"two\" class=\"hidden peer\" @change=\"$refs.custominput.value =0;total = 200 \"> <label for=\"two\" class=\"grid place-content-center rounded-full border border-blue-400 bg-blue-50 peer-checked:bg-blue-500 peer-checked:text-white h-10 w-10\">2</label></div><div><input type=\"radio\" value=\"3\" name=\"amount\" id=\"three\" class=\"hidden peer\" @change=\"$refs.custominput.value =0;total = 300 \"> <label for=\"three\" class=\"grid place-content-center rounded-full border border-blue-400 bg-blue-50 peer-checked:bg-blue-500 peer-checked:text-white h-10 w-10\">3</label></div><div class=\"border border-blue-400 bg-blue-50 rounded-xl overflow-clip\"><input type=\"text\" x-ref=\"custominput\" name=\"custom\" id=\"custom-amount\" class=\"peer-checked:bg-blue-500 peer-checked:text-white\n\t\t\t\t\t\t\t\t h-10 w-10 text-center\" x-mask=\"9999\" placeholder=\"10\" @input=\"$refs.custom.checked = true; total = $event.target.value*100\" value=\"0\"> <input type=\"radio\" x-ref=\"custom\" value=\"custom\" name=\"amount\" id=\"custom\" class=\"hidden peer\"></div></div><input type=\"text\" name=\"name\" id=\"name\" placeholder=\"Your name\" class=\"p-2 border border-gray-300 rounded-xl\" required max=\"100\"> <input type=\"email\" name=\"email\" id=\"email\" placeholder=\"Your email\" class=\"p-2 border border-gray-300 rounded-xl\" required max=\"100\"> <textarea name=\"message\" id=\"message\" rows=\"3\" class=\"p-2 border border-gray-300 rounded-xl\" placeholder=\"A very nice message for me\" required maxlength=\"500\"></textarea> <button type=\"submit\" class=\"bg-blue-500 text-white p-2 rounded-xl\">NRS <span x-text=\"total\"></span></button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
