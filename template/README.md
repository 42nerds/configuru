# README Test

My age is {{ .Answers.age.Value }}.
{{- if .Answers.gender }}
My gender is {{ .Answers.gender.Value }}.
{{- end }}
