### Fastener thread data table
id: integer, primary key
text_key: text
series: text
class: text
basic_diam: real
n_thd: integer
pitch: real
dmaj_min: real
dmaj_max: real
dpitch_min: real
dpitch_max: real
dmin_min: real
dmin_max: real


text_key is derived from basic_diam, series, and thread class
example: 0.2500_UNJF_3A, 0.190_UNC_2A