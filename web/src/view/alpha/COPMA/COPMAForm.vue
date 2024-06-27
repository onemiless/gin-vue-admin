<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="COMPANY字段:" prop="COMPANY">
          <el-input v-model="formData.COMPANY" :clearable="true"  placeholder="请输入COMPANY字段" />
       </el-form-item>
        <el-form-item label="CREATOR字段:" prop="CREATOR">
          <el-input v-model="formData.CREATOR" :clearable="true"  placeholder="请输入CREATOR字段" />
       </el-form-item>
        <el-form-item label="USRGROUP字段:" prop="USRGROUP">
          <el-input v-model="formData.USRGROUP" :clearable="true"  placeholder="请输入USRGROUP字段" />
       </el-form-item>
        <el-form-item label="CREATEDATE字段:" prop="CREATEDATE">
          <el-input v-model="formData.CREATEDATE" :clearable="true"  placeholder="请输入CREATEDATE字段" />
       </el-form-item>
        <el-form-item label="MODIFIER字段:" prop="MODIFIER">
          <el-input v-model="formData.MODIFIER" :clearable="true"  placeholder="请输入MODIFIER字段" />
       </el-form-item>
        <el-form-item label="MODIDATE字段:" prop="MODIDATE">
          <el-input v-model="formData.MODIDATE" :clearable="true"  placeholder="请输入MODIDATE字段" />
       </el-form-item>
        <el-form-item label="FLAG字段:" prop="FLAG">
          <el-input-number v-model="formData.FLAG" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA001字段:" prop="MA001">
          <el-input v-model="formData.MA001" :clearable="true"  placeholder="请输入MA001字段" />
       </el-form-item>
        <el-form-item label="MA002字段:" prop="MA002">
          <el-input v-model="formData.MA002" :clearable="true"  placeholder="请输入MA002字段" />
       </el-form-item>
        <el-form-item label="MA003字段:" prop="MA003">
          <el-input v-model="formData.MA003" :clearable="true"  placeholder="请输入MA003字段" />
       </el-form-item>
        <el-form-item label="MA004字段:" prop="MA004">
          <el-input v-model="formData.MA004" :clearable="true"  placeholder="请输入MA004字段" />
       </el-form-item>
        <el-form-item label="MA005字段:" prop="MA005">
          <el-input v-model="formData.MA005" :clearable="true"  placeholder="请输入MA005字段" />
       </el-form-item>
        <el-form-item label="MA006字段:" prop="MA006">
          <el-input v-model="formData.MA006" :clearable="true"  placeholder="请输入MA006字段" />
       </el-form-item>
        <el-form-item label="MA007字段:" prop="MA007">
          <el-input v-model="formData.MA007" :clearable="true"  placeholder="请输入MA007字段" />
       </el-form-item>
        <el-form-item label="MA008字段:" prop="MA008">
          <el-input v-model="formData.MA008" :clearable="true"  placeholder="请输入MA008字段" />
       </el-form-item>
        <el-form-item label="MA009字段:" prop="MA009">
          <el-input v-model="formData.MA009" :clearable="true"  placeholder="请输入MA009字段" />
       </el-form-item>
        <el-form-item label="MA010字段:" prop="MA010">
          <el-input v-model="formData.MA010" :clearable="true"  placeholder="请输入MA010字段" />
       </el-form-item>
        <el-form-item label="MA011字段:" prop="MA011">
          <el-input-number v-model="formData.MA011" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA012字段:" prop="MA012">
          <el-input-number v-model="formData.MA012" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA013字段:" prop="MA013">
          <el-input-number v-model="formData.MA013" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA014字段:" prop="MA014">
          <el-input v-model="formData.MA014" :clearable="true"  placeholder="请输入MA014字段" />
       </el-form-item>
        <el-form-item label="MA015字段:" prop="MA015">
          <el-input v-model="formData.MA015" :clearable="true"  placeholder="请输入MA015字段" />
       </el-form-item>
        <el-form-item label="MA016字段:" prop="MA016">
          <el-input v-model="formData.MA016" :clearable="true"  placeholder="请输入MA016字段" />
       </el-form-item>
        <el-form-item label="MA017字段:" prop="MA017">
          <el-input v-model="formData.MA017" :clearable="true"  placeholder="请输入MA017字段" />
       </el-form-item>
        <el-form-item label="MA018字段:" prop="MA018">
          <el-input v-model="formData.MA018" :clearable="true"  placeholder="请输入MA018字段" />
       </el-form-item>
        <el-form-item label="MA019字段:" prop="MA019">
          <el-input v-model="formData.MA019" :clearable="true"  placeholder="请输入MA019字段" />
       </el-form-item>
        <el-form-item label="MA020字段:" prop="MA020">
          <el-input v-model="formData.MA020" :clearable="true"  placeholder="请输入MA020字段" />
       </el-form-item>
        <el-form-item label="MA021字段:" prop="MA021">
          <el-input v-model="formData.MA021" :clearable="true"  placeholder="请输入MA021字段" />
       </el-form-item>
        <el-form-item label="MA022字段:" prop="MA022">
          <el-input v-model="formData.MA022" :clearable="true"  placeholder="请输入MA022字段" />
       </el-form-item>
        <el-form-item label="MA023字段:" prop="MA023">
          <el-input v-model="formData.MA023" :clearable="true"  placeholder="请输入MA023字段" />
       </el-form-item>
        <el-form-item label="MA024字段:" prop="MA024">
          <el-input v-model="formData.MA024" :clearable="true"  placeholder="请输入MA024字段" />
       </el-form-item>
        <el-form-item label="MA025字段:" prop="MA025">
          <el-input v-model="formData.MA025" :clearable="true"  placeholder="请输入MA025字段" />
       </el-form-item>
        <el-form-item label="MA026字段:" prop="MA026">
          <el-input v-model="formData.MA026" :clearable="true"  placeholder="请输入MA026字段" />
       </el-form-item>
        <el-form-item label="MA027字段:" prop="MA027">
          <el-input v-model="formData.MA027" :clearable="true"  placeholder="请输入MA027字段" />
       </el-form-item>
        <el-form-item label="MA028字段:" prop="MA028">
          <el-input v-model="formData.MA028" :clearable="true"  placeholder="请输入MA028字段" />
       </el-form-item>
        <el-form-item label="MA029字段:" prop="MA029">
          <el-input v-model="formData.MA029" :clearable="true"  placeholder="请输入MA029字段" />
       </el-form-item>
        <el-form-item label="MA030字段:" prop="MA030">
          <el-input v-model="formData.MA030" :clearable="true"  placeholder="请输入MA030字段" />
       </el-form-item>
        <el-form-item label="MA031字段:" prop="MA031">
          <el-input v-model="formData.MA031" :clearable="true"  placeholder="请输入MA031字段" />
       </el-form-item>
        <el-form-item label="MA032字段:" prop="MA032">
          <el-input v-model="formData.MA032" :clearable="true"  placeholder="请输入MA032字段" />
       </el-form-item>
        <el-form-item label="MA033字段:" prop="MA033">
          <el-input-number v-model="formData.MA033" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA034字段:" prop="MA034">
          <el-input-number v-model="formData.MA034" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA035字段:" prop="MA035">
          <el-input v-model="formData.MA035" :clearable="true"  placeholder="请输入MA035字段" />
       </el-form-item>
        <el-form-item label="MA036字段:" prop="MA036">
          <el-input-number v-model="formData.MA036" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA037字段:" prop="MA037">
          <el-input v-model="formData.MA037" :clearable="true"  placeholder="请输入MA037字段" />
       </el-form-item>
        <el-form-item label="MA038字段:" prop="MA038">
          <el-input v-model="formData.MA038" :clearable="true"  placeholder="请输入MA038字段" />
       </el-form-item>
        <el-form-item label="MA039字段:" prop="MA039">
          <el-input v-model="formData.MA039" :clearable="true"  placeholder="请输入MA039字段" />
       </el-form-item>
        <el-form-item label="MA040字段:" prop="MA040">
          <el-input v-model="formData.MA040" :clearable="true"  placeholder="请输入MA040字段" />
       </el-form-item>
        <el-form-item label="MA041字段:" prop="MA041">
          <el-input v-model="formData.MA041" :clearable="true"  placeholder="请输入MA041字段" />
       </el-form-item>
        <el-form-item label="MA042字段:" prop="MA042">
          <el-input v-model="formData.MA042" :clearable="true"  placeholder="请输入MA042字段" />
       </el-form-item>
        <el-form-item label="MA043字段:" prop="MA043">
          <el-input v-model="formData.MA043" :clearable="true"  placeholder="请输入MA043字段" />
       </el-form-item>
        <el-form-item label="MA044字段:" prop="MA044">
          <el-input-number v-model="formData.MA044" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA045字段:" prop="MA045">
          <el-input-number v-model="formData.MA045" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA046字段:" prop="MA046">
          <el-input v-model="formData.MA046" :clearable="true"  placeholder="请输入MA046字段" />
       </el-form-item>
        <el-form-item label="MA047字段:" prop="MA047">
          <el-input v-model="formData.MA047" :clearable="true"  placeholder="请输入MA047字段" />
       </el-form-item>
        <el-form-item label="MA048字段:" prop="MA048">
          <el-input v-model="formData.MA048" :clearable="true"  placeholder="请输入MA048字段" />
       </el-form-item>
        <el-form-item label="MA049字段:" prop="MA049">
          <el-input v-model="formData.MA049" :clearable="true"  placeholder="请输入MA049字段" />
       </el-form-item>
        <el-form-item label="MA050字段:" prop="MA050">
          <el-input v-model="formData.MA050" :clearable="true"  placeholder="请输入MA050字段" />
       </el-form-item>
        <el-form-item label="MA051字段:" prop="MA051">
          <el-input v-model="formData.MA051" :clearable="true"  placeholder="请输入MA051字段" />
       </el-form-item>
        <el-form-item label="MA052字段:" prop="MA052">
          <el-input v-model="formData.MA052" :clearable="true"  placeholder="请输入MA052字段" />
       </el-form-item>
        <el-form-item label="MA053字段:" prop="MA053">
          <el-input v-model="formData.MA053" :clearable="true"  placeholder="请输入MA053字段" />
       </el-form-item>
        <el-form-item label="MA054字段:" prop="MA054">
          <el-input v-model="formData.MA054" :clearable="true"  placeholder="请输入MA054字段" />
       </el-form-item>
        <el-form-item label="MA055字段:" prop="MA055">
          <el-input v-model="formData.MA055" :clearable="true"  placeholder="请输入MA055字段" />
       </el-form-item>
        <el-form-item label="MA056字段:" prop="MA056">
          <el-input v-model="formData.MA056" :clearable="true"  placeholder="请输入MA056字段" />
       </el-form-item>
        <el-form-item label="MA057字段:" prop="MA057">
          <el-input v-model="formData.MA057" :clearable="true"  placeholder="请输入MA057字段" />
       </el-form-item>
        <el-form-item label="MA058字段:" prop="MA058">
          <el-input v-model="formData.MA058" :clearable="true"  placeholder="请输入MA058字段" />
       </el-form-item>
        <el-form-item label="MA059字段:" prop="MA059">
          <el-input-number v-model="formData.MA059" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA060字段:" prop="MA060">
          <el-input-number v-model="formData.MA060" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA061字段:" prop="MA061">
          <el-input-number v-model="formData.MA061" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA062字段:" prop="MA062">
          <el-input v-model="formData.MA062" :clearable="true"  placeholder="请输入MA062字段" />
       </el-form-item>
        <el-form-item label="MA063字段:" prop="MA063">
          <el-input v-model="formData.MA063" :clearable="true"  placeholder="请输入MA063字段" />
       </el-form-item>
        <el-form-item label="MA064字段:" prop="MA064">
          <el-input v-model="formData.MA064" :clearable="true"  placeholder="请输入MA064字段" />
       </el-form-item>
        <el-form-item label="MA065字段:" prop="MA065">
          <el-input v-model="formData.MA065" :clearable="true"  placeholder="请输入MA065字段" />
       </el-form-item>
        <el-form-item label="MA066字段:" prop="MA066">
          <el-input v-model="formData.MA066" :clearable="true"  placeholder="请输入MA066字段" />
       </el-form-item>
        <el-form-item label="MA067字段:" prop="MA067">
          <el-input-number v-model="formData.MA067" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA068字段:" prop="MA068">
          <el-input v-model="formData.MA068" :clearable="true"  placeholder="请输入MA068字段" />
       </el-form-item>
        <el-form-item label="MA069字段:" prop="MA069">
          <el-input v-model="formData.MA069" :clearable="true"  placeholder="请输入MA069字段" />
       </el-form-item>
        <el-form-item label="MA070字段:" prop="MA070">
          <el-input v-model="formData.MA070" :clearable="true"  placeholder="请输入MA070字段" />
       </el-form-item>
        <el-form-item label="MA071字段:" prop="MA071">
          <el-input v-model="formData.MA071" :clearable="true"  placeholder="请输入MA071字段" />
       </el-form-item>
        <el-form-item label="MA072字段:" prop="MA072">
          <el-input v-model="formData.MA072" :clearable="true"  placeholder="请输入MA072字段" />
       </el-form-item>
        <el-form-item label="MA073字段:" prop="MA073">
          <el-input v-model="formData.MA073" :clearable="true"  placeholder="请输入MA073字段" />
       </el-form-item>
        <el-form-item label="MA074字段:" prop="MA074">
          <el-input v-model="formData.MA074" :clearable="true"  placeholder="请输入MA074字段" />
       </el-form-item>
        <el-form-item label="MA075字段:" prop="MA075">
          <el-input v-model="formData.MA075" :clearable="true"  placeholder="请输入MA075字段" />
       </el-form-item>
        <el-form-item label="MA076字段:" prop="MA076">
          <el-input v-model="formData.MA076" :clearable="true"  placeholder="请输入MA076字段" />
       </el-form-item>
        <el-form-item label="MA077字段:" prop="MA077">
          <el-input v-model="formData.MA077" :clearable="true"  placeholder="请输入MA077字段" />
       </el-form-item>
        <el-form-item label="MA078字段:" prop="MA078">
          <el-input v-model="formData.MA078" :clearable="true"  placeholder="请输入MA078字段" />
       </el-form-item>
        <el-form-item label="MA079字段:" prop="MA079">
          <el-input v-model="formData.MA079" :clearable="true"  placeholder="请输入MA079字段" />
       </el-form-item>
        <el-form-item label="MA080字段:" prop="MA080">
          <el-input v-model="formData.MA080" :clearable="true"  placeholder="请输入MA080字段" />
       </el-form-item>
        <el-form-item label="MA081字段:" prop="MA081">
          <el-input v-model="formData.MA081" :clearable="true"  placeholder="请输入MA081字段" />
       </el-form-item>
        <el-form-item label="MA082字段:" prop="MA082">
          <el-input v-model="formData.MA082" :clearable="true"  placeholder="请输入MA082字段" />
       </el-form-item>
        <el-form-item label="MA083字段:" prop="MA083">
          <el-input v-model="formData.MA083" :clearable="true"  placeholder="请输入MA083字段" />
       </el-form-item>
        <el-form-item label="MA084字段:" prop="MA084">
          <el-input v-model="formData.MA084" :clearable="true"  placeholder="请输入MA084字段" />
       </el-form-item>
        <el-form-item label="MA085字段:" prop="MA085">
          <el-input v-model="formData.MA085" :clearable="true"  placeholder="请输入MA085字段" />
       </el-form-item>
        <el-form-item label="MA086字段:" prop="MA086">
          <el-input v-model="formData.MA086" :clearable="true"  placeholder="请输入MA086字段" />
       </el-form-item>
        <el-form-item label="MA087字段:" prop="MA087">
          <el-input v-model="formData.MA087" :clearable="true"  placeholder="请输入MA087字段" />
       </el-form-item>
        <el-form-item label="MA088字段:" prop="MA088">
          <el-input v-model="formData.MA088" :clearable="true"  placeholder="请输入MA088字段" />
       </el-form-item>
        <el-form-item label="MA089字段:" prop="MA089">
          <el-input v-model="formData.MA089" :clearable="true"  placeholder="请输入MA089字段" />
       </el-form-item>
        <el-form-item label="MA090字段:" prop="MA090">
          <el-input-number v-model="formData.MA090" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA091字段:" prop="MA091">
          <el-input-number v-model="formData.MA091" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA092字段:" prop="MA092">
          <el-input-number v-model="formData.MA092" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA093字段:" prop="MA093">
          <el-input-number v-model="formData.MA093" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA094字段:" prop="MA094">
          <el-input-number v-model="formData.MA094" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA095字段:" prop="MA095">
          <el-input-number v-model="formData.MA095" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA096字段:" prop="MA096">
          <el-input v-model="formData.MA096" :clearable="true"  placeholder="请输入MA096字段" />
       </el-form-item>
        <el-form-item label="MA097字段:" prop="MA097">
          <el-input v-model="formData.MA097" :clearable="true"  placeholder="请输入MA097字段" />
       </el-form-item>
        <el-form-item label="MA098字段:" prop="MA098">
          <el-input v-model="formData.MA098" :clearable="true"  placeholder="请输入MA098字段" />
       </el-form-item>
        <el-form-item label="MA099字段:" prop="MA099">
          <el-input v-model="formData.MA099" :clearable="true"  placeholder="请输入MA099字段" />
       </el-form-item>
        <el-form-item label="MA100字段:" prop="MA100">
          <el-input v-model="formData.MA100" :clearable="true"  placeholder="请输入MA100字段" />
       </el-form-item>
        <el-form-item label="MA101字段:" prop="MA101">
          <el-input-number v-model="formData.MA101" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA102字段:" prop="MA102">
          <el-input v-model="formData.MA102" :clearable="true"  placeholder="请输入MA102字段" />
       </el-form-item>
        <el-form-item label="MA103字段:" prop="MA103">
          <el-input v-model="formData.MA103" :clearable="true"  placeholder="请输入MA103字段" />
       </el-form-item>
        <el-form-item label="MA104字段:" prop="MA104">
          <el-input-number v-model="formData.MA104" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA105字段:" prop="MA105">
          <el-input v-model="formData.MA105" :clearable="true"  placeholder="请输入MA105字段" />
       </el-form-item>
        <el-form-item label="MA106字段:" prop="MA106">
          <el-input v-model="formData.MA106" :clearable="true"  placeholder="请输入MA106字段" />
       </el-form-item>
        <el-form-item label="MA107字段:" prop="MA107">
          <el-input v-model="formData.MA107" :clearable="true"  placeholder="请输入MA107字段" />
       </el-form-item>
        <el-form-item label="MA108字段:" prop="MA108">
          <el-input v-model="formData.MA108" :clearable="true"  placeholder="请输入MA108字段" />
       </el-form-item>
        <el-form-item label="MA109字段:" prop="MA109">
          <el-input-number v-model="formData.MA109" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA110字段:" prop="MA110">
          <el-input-number v-model="formData.MA110" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA111字段:" prop="MA111">
          <el-input-number v-model="formData.MA111" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA112字段:" prop="MA112">
          <el-input v-model="formData.MA112" :clearable="true"  placeholder="请输入MA112字段" />
       </el-form-item>
        <el-form-item label="MA113字段:" prop="MA113">
          <el-input v-model="formData.MA113" :clearable="true"  placeholder="请输入MA113字段" />
       </el-form-item>
        <el-form-item label="MA114字段:" prop="MA114">
          <el-input v-model="formData.MA114" :clearable="true"  placeholder="请输入MA114字段" />
       </el-form-item>
        <el-form-item label="MAB01字段:" prop="MAB01">
          <el-input v-model="formData.MAB01" :clearable="true"  placeholder="请输入MAB01字段" />
       </el-form-item>
        <el-form-item label="MA115字段:" prop="MA115">
          <el-input v-model="formData.MA115" :clearable="true"  placeholder="请输入MA115字段" />
       </el-form-item>
        <el-form-item label="UDF01字段:" prop="UDF01">
          <el-input v-model="formData.UDF01" :clearable="true"  placeholder="请输入UDF01字段" />
       </el-form-item>
        <el-form-item label="UDF02字段:" prop="UDF02">
          <el-input v-model="formData.UDF02" :clearable="true"  placeholder="请输入UDF02字段" />
       </el-form-item>
        <el-form-item label="UDF03字段:" prop="UDF03">
          <el-input v-model="formData.UDF03" :clearable="true"  placeholder="请输入UDF03字段" />
       </el-form-item>
        <el-form-item label="UDF04字段:" prop="UDF04">
          <el-input v-model="formData.UDF04" :clearable="true"  placeholder="请输入UDF04字段" />
       </el-form-item>
        <el-form-item label="UDF05字段:" prop="UDF05">
          <el-input v-model="formData.UDF05" :clearable="true"  placeholder="请输入UDF05字段" />
       </el-form-item>
        <el-form-item label="UDF06字段:" prop="UDF06">
          <el-input v-model="formData.UDF06" :clearable="true"  placeholder="请输入UDF06字段" />
       </el-form-item>
        <el-form-item label="UDF51字段:" prop="UDF51">
          <el-input-number v-model="formData.UDF51" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF52字段:" prop="UDF52">
          <el-input-number v-model="formData.UDF52" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF53字段:" prop="UDF53">
          <el-input-number v-model="formData.UDF53" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF54字段:" prop="UDF54">
          <el-input-number v-model="formData.UDF54" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF55字段:" prop="UDF55">
          <el-input-number v-model="formData.UDF55" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF56字段:" prop="UDF56">
          <el-input-number v-model="formData.UDF56" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF07字段:" prop="UDF07">
          <el-input v-model="formData.UDF07" :clearable="true"  placeholder="请输入UDF07字段" />
       </el-form-item>
        <el-form-item label="UDF08字段:" prop="UDF08">
          <el-input v-model="formData.UDF08" :clearable="true"  placeholder="请输入UDF08字段" />
       </el-form-item>
        <el-form-item label="UDF09字段:" prop="UDF09">
          <el-input v-model="formData.UDF09" :clearable="true"  placeholder="请输入UDF09字段" />
       </el-form-item>
        <el-form-item label="UDF10字段:" prop="UDF10">
          <el-input v-model="formData.UDF10" :clearable="true"  placeholder="请输入UDF10字段" />
       </el-form-item>
        <el-form-item label="UDF11字段:" prop="UDF11">
          <el-input v-model="formData.UDF11" :clearable="true"  placeholder="请输入UDF11字段" />
       </el-form-item>
        <el-form-item label="UDF12字段:" prop="UDF12">
          <el-input v-model="formData.UDF12" :clearable="true"  placeholder="请输入UDF12字段" />
       </el-form-item>
        <el-form-item label="UDF57字段:" prop="UDF57">
          <el-input-number v-model="formData.UDF57" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF58字段:" prop="UDF58">
          <el-input-number v-model="formData.UDF58" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF59字段:" prop="UDF59">
          <el-input-number v-model="formData.UDF59" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF60字段:" prop="UDF60">
          <el-input-number v-model="formData.UDF60" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF61字段:" prop="UDF61">
          <el-input-number v-model="formData.UDF61" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="UDF62字段:" prop="UDF62">
          <el-input-number v-model="formData.UDF62" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="MA200字段:" prop="MA200">
          <el-input-number v-model="formData.MA200" :precision="3" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createCOPMA,
  updateCOPMA,
  findCOPMA
} from '@/api/alpha/COPMA'

defineOptions({
    name: 'COPMAForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            COMPANY: '',
            CREATOR: '',
            USRGROUP: '',
            CREATEDATE: '',
            MODIFIER: '',
            MODIDATE: '',
            FLAG: 0,
            MA001: '',
            MA002: '',
            MA003: '',
            MA004: '',
            MA005: '',
            MA006: '',
            MA007: '',
            MA008: '',
            MA009: '',
            MA010: '',
            MA011: 0,
            MA012: 0,
            MA013: 0,
            MA014: '',
            MA015: '',
            MA016: '',
            MA017: '',
            MA018: '',
            MA019: '',
            MA020: '',
            MA021: '',
            MA022: '',
            MA023: '',
            MA024: '',
            MA025: '',
            MA026: '',
            MA027: '',
            MA028: '',
            MA029: '',
            MA030: '',
            MA031: '',
            MA032: '',
            MA033: 0,
            MA034: 0,
            MA035: '',
            MA036: 0,
            MA037: '',
            MA038: '',
            MA039: '',
            MA040: '',
            MA041: '',
            MA042: '',
            MA043: '',
            MA044: 0,
            MA045: 0,
            MA046: '',
            MA047: '',
            MA048: '',
            MA049: '',
            MA050: '',
            MA051: '',
            MA052: '',
            MA053: '',
            MA054: '',
            MA055: '',
            MA056: '',
            MA057: '',
            MA058: '',
            MA059: 0,
            MA060: 0,
            MA061: 0,
            MA062: '',
            MA063: '',
            MA064: '',
            MA065: '',
            MA066: '',
            MA067: 0,
            MA068: '',
            MA069: '',
            MA070: '',
            MA071: '',
            MA072: '',
            MA073: '',
            MA074: '',
            MA075: '',
            MA076: '',
            MA077: '',
            MA078: '',
            MA079: '',
            MA080: '',
            MA081: '',
            MA082: '',
            MA083: '',
            MA084: '',
            MA085: '',
            MA086: '',
            MA087: '',
            MA088: '',
            MA089: '',
            MA090: 0,
            MA091: 0,
            MA092: 0,
            MA093: 0,
            MA094: 0,
            MA095: 0,
            MA096: '',
            MA097: '',
            MA098: '',
            MA099: '',
            MA100: '',
            MA101: 0,
            MA102: '',
            MA103: '',
            MA104: 0,
            MA105: '',
            MA106: '',
            MA107: '',
            MA108: '',
            MA109: 0,
            MA110: 0,
            MA111: 0,
            MA112: '',
            MA113: '',
            MA114: '',
            MAB01: '',
            MA115: '',
            UDF01: '',
            UDF02: '',
            UDF03: '',
            UDF04: '',
            UDF05: '',
            UDF06: '',
            UDF51: 0,
            UDF52: 0,
            UDF53: 0,
            UDF54: 0,
            UDF55: 0,
            UDF56: 0,
            UDF07: '',
            UDF08: '',
            UDF09: '',
            UDF10: '',
            UDF11: '',
            UDF12: '',
            UDF57: 0,
            UDF58: 0,
            UDF59: 0,
            UDF60: 0,
            UDF61: 0,
            UDF62: 0,
            MA200: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCOPMA({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recopma
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCOPMA(formData.value)
               break
             case 'update':
               res = await updateCOPMA(formData.value)
               break
             default:
               res = await createCOPMA(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
