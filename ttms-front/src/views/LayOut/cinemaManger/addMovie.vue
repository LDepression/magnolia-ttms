<script setup>
//剧目管理
import { ArrowRight } from "@element-plus/icons-vue";
import { onMounted, onUnmounted, reactive } from "vue";
import { fileUpLoad, changeAvatar } from "../../../apis/user";
import { addMovie } from "../../../apis/movie";
import axios from "axios";
import { ElMessage } from "element-plus";
// import { useStore, useMutations } from 'vuex'
// const store = useStore()
// const {getMovieId} = useMutations({
//   getMovieId:'getMovieId'
// })

const form = reactive({
  area: "",
  avatar: "",
  content: "",
  director: "",
  duration: 1,
  actors: [""],
  name: "",
  show_time: "",
  tags: [""],
});
const changeTime = () => {
  const timeStr = form.show_time;
  const dateObj = new Date(timeStr);
  const timestamp = Math.floor(dateObj.getTime() / 1000);
  form.show_time = timestamp;
  console.log(form.show_time);
};
const holdMovieId = (movieId) => {
  getMovieId(movieId);
};
const handleFileChange = async () => {
  const File = event.target.files[0];
  console.log(File);
  //构造format对象
  const formData = new FormData();
  formData.append("File", File);
  try {
    //发送请求上传图片
    const response = await fileUpLoad(formData);

    form.avatar = response.data.data.URL;
    console.log(form.avatar);
    // const AvatarUrl = await changeAvatar(form.avatar);
  } catch (error) {
    console.log(error);
  }
};

const createMovie = async () => {
  
  changeTime()
  let x = form.actors.toString().split(",")
  form.actors = x
  console.log(form.actors);
  let y = form.tags.toString().split(",")
  form.tags = y
  let res = await addMovie({
    actors: form.actors,
    area: form.area,
    avatar: form.avatar,
    content: form.content,
    director: form.director,
    duration: Number(form.duration),
    name: form.name,
    show_time: form.show_time,
    tags: form.tags,
  });

  if (res.data.status_code == 0) {
    ElMessage({
      message: "添加成功!!",
      center: true,
      type: "success",
    });
    const movieId = res.data.data.movieID;
   
  }
};
// onMounted(()=>{sendData()})
// onUnmounted(()=>{holdMovieId(movieId)})
</script>
<template>
  <div class="bread">
    <el-breadcrumb :separator-icon="ArrowRight">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item :to="{}">剧院管理</el-breadcrumb-item>
      <el-breadcrumb-item :to="{}">电影信息管理</el-breadcrumb-item>
    </el-breadcrumb>
    <hr />
  </div>
  <div class="form">
    <el-form :model="form" label-width="120px">
      <el-form-item label="请输入电影名称">
        <el-input style="width: 200px" v-model="form.name" />
      </el-form-item>
      <el-form-item label="请输入导演名称">
        <el-input style="width: 200px" v-model="form.director" />
      </el-form-item>
      <el-form-item label="请输入电影时长">
        <el-input style="width: 200px" v-model="form.duration" />
      </el-form-item>
      <el-form-item label="地区">
        <el-select v-model="form.area" placeholder="请选择电影地区">
          <el-option value="大陆" />
          <el-option value="美国" />
          <el-option value="韩国" />
          <el-option value="日本" />
          <el-option value="中国香港" />
          <el-option value="中国台湾" />
          <el-option value="泰国" />
          <el-option value="印度" />
          <el-option value="法国" />
          <el-option value="英国" />
          <el-option value="俄罗斯" />
          <el-option value="意大利" />
          <el-option value="西班牙" />
          <el-option value="德国" />
          <el-option value="波兰" />
          <el-option value="澳大利亚" />
          <el-option value="伊朗" />
          <el-option value="其他" />
        </el-select>
      </el-form-item>
      <el-form-item label="电影类型">
        <el-select v-model="form.tags" placeholder="请选择电影类型">
          <el-option value="爱情" />
          <el-option value="喜剧" />
          <el-option value="动画" />
          <el-option value="剧情" />
          <el-option value="恐怖" />
          <el-option value="惊悚" />
          <el-option value="科幻" />
          <el-option value="动作" />
          <el-option value="悬疑" />
          <el-option value="犯罪" />
          <el-option value="冒险" />
          <el-option value="战争" />
          <el-option value="奇幻" />
          <el-option value="运动" />
          <el-option value="家庭" />
          <el-option value="古装" />
          <el-option value="武侠" />
          <el-option value="西部" />
          <el-option value="历史" />
          <el-option value="传记" />
          <el-option value="歌舞" />
          <el-option value="黑色电影" />
          <el-option value="短片" />
          <el-option value="纪录片" />
          <el-option value="戏曲" />
          <el-option value="音乐" />
          <el-option value="灾难" />
          <el-option value="青春" />
          <el-option value="儿童" />
          <el-option value="其他" />
        </el-select>
      </el-form-item>
      <el-form-item label="上映时间">
        <el-col :span="11">
          <el-date-picker
            v-model="form.show_time"
            type="date"
            placeholder="Pick a date"
            style="width: 100%"
          />
        </el-col>
        <el-col :span="2" class="text-center">
          <span class="text-gray-500">-</span>
        </el-col>
        <el-col :span="11">
          <el-time-picker
            v-model="form.show_time"
            placeholder="Pick a time"
            style="width: 100%"
          />
        </el-col>
      </el-form-item>
      <el-form-item label="电影主要演员">
        <el-input v-model="form.actors" type="textarea" />
      </el-form-item>

      <el-form-item label="电影大致内容">
        <el-input v-model="form.content" type="textarea" />
      </el-form-item>
      <el-form-item label="上传头像">
        <input
          type="file"
          class="file"
          name="File"
          @change="handleFileChange"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="createMovie">Create</el-button>
        <el-button>Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<style lang="css" scoped>
.bread {
  color: #fff;
}
</style>