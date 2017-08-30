{% extends 'full.tpl'%}

{% block in_prompt %}
{% endblock in_prompt %}

{% block input_group %}
{{ super() }}
{% endblock input_group %}

{% block output_area_prompt %}
{% endblock output_area_prompt %}

{% block error %}
{% endblock error %}

{% block stream_stderr %}
{% endblock stream_stderr %}

{% block stream_stdout %}
{% endblock stream_stdout %}