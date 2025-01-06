# Open-Source Library Licenses

Gravwell incorporates compatibly-licensed open-source code. This page links to compilations of licenses for software used in the backend and the web frontend.

% Use an "a" tag here, since sphinx can't integrity-check the license file links

```{datatemplate:json} open_source.json
{% for license in data %}
- <a href="{{ license.name }}">{{ license.name }}</a>
{% endfor %}
```
