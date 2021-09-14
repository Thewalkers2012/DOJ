<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
      <router-link to="/" class="navbar-brand">DOJ</router-link>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" href="#">题库</a>
          </li>
          <li class="nav-item">
            <router-link to="/about" class="nav-link">关于我们</router-link>
          </li>
        </ul>
        <div class="row">
          <button
            class="btn btn-outline-primary col"
            @click="this.$router.push('/login')"
          >
            登录
          </button>
          <button
            class="btn btn-outline-success col"
            @click="this.$router.push('/register')"
          >
            注册
          </button>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'DojNavbar',
}
</script>