package release

var template string = `
<% if origin -%>
Origin: <%= origin %>
<% end -%>
Codename: {{ codename }}
Date: {{ date }}
Architectures: {{ architectures }}
Components: {{ components }}
Suite: {{ suite }}
MD5Sum:
{{ #packages }} 
<%= p[:md5] %> <%= p[:size].to_s.rjust(16) %> <%= f %>
{{ /packages }}
SHA1:
{{ #packages }} 
<%= p[:sha1] %> <%= p[:size].to_s.rjust(16) %> <%= f %>
{{ /packages }}
SHA256:
{{ #packages }}
 <%= p[:sha256] %> <%= p[:size].to_s.rjust(16) %> <%= f %>
{{ /packages }}
`

func getTemplate() string {
	return template
}
