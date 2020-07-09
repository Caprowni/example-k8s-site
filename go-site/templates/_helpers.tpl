{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "name" -}}
{{- default .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
The release labels for objects
*/}}
{{- define "releaseLabels" }}
app: "go-test"
chart: "{{ .Chart.Name }}-{{ .Chart.Version }}"
release: "{{ .Release.Name }}"
heritage: "{{ .Release.Service }}"
image_tag: "{{ .Values.image.tag }}"
{{- end }}

{{/*
The selection labels for objects
*/}}
{{- define "selectionLabels" }}
app: "go-test"
release: "{{ .Release.Name }}"
{{- end }}
