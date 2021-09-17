<template>
  <div class="Problem">
    <b-card class="mt-3">
      <h2 class="mt-3" style="font-weight: bold">
        {{ problem.problem_id }}. {{ problem.problem_name }}
      </h2>
      <div class="mt-3">
        <b-nav tabs fill style="color: black">
          <b-nav-item>
            <div style="color: black">题目</div>
          </b-nav-item>
          <b-nav-item>
            <div style="color: black">提交记录</div>
          </b-nav-item>
          <b-nav-item>
            <div style="color: black">讨论</div>
          </b-nav-item>
        </b-nav>
      </div>
      <div class="mt-3">
        <h3 style="font-weight: bold">题目描述：</h3>
        <div style="font-size: 20px" class="mt-3">
          {{ problem.description }}
        </div>
      </div>
      <div class="mt-3">
        <h3 style="font-weight: bold">输入格式：</h3>
        <div style="font-size: 20px" class="mt-3">
          {{ problem.test_case }}
        </div>
      </div>
      <div style="font-size: 20px" class="mt-3">
        <h3 style="font-weight: bold">输出格式：</h3>
        {{ problem.test_case }}
      </div>
      <!-- 代码编辑区域 -->
      <b-card class="mt-5" title="编辑器">
        <pre v-highlightjs><code class="cpp">{{ code }}</code></pre>
      </b-card>
    </b-card>
  </div>
</template>

<script>
import problemService from '../service/problemService';

export default {
  data() {
    return {
      problem: {},
      code: `
#include <bits/stdc++.h>
using namespace std;

const int MAX_N = 100010;

int h[MAX_N], e[MAX_N], w[MAX_N], ne[MAX_N], idx;

void add(int a, int b, int c) {
  e[idx] = b, w[idx] = c, ne[idx] = h[a], h[a] = idx++;
}

int n, m, s;
vector<pair<int, int>> adj[MAX_N];
long long dist[MAX_N];

typedef pair<int, int> pii;

void dijkstra() {
  priority_queue<pii, vector<pii>, greater<pii>> heap;
  vector<bool> vis(n);
  heap.push(make_pair(0, s));
  memset(dist, 0x3f, sizeof dist);
  dist[s] = 0;
  while (!heap.empty()) {
    auto t = heap.top();
    heap.pop();
    int ver = t.second;
    if (vis[ver]) {
      continue;
    }
    vis[ver] = true;
    for (auto &v : adj[ver]) {
      if (dist[v.first] > dist[ver] + v.second) {
        dist[v.first] = dist[ver] + v.second;
        heap.push({dist[v.first], v.first});
      }
    }
  }
  for (int i = 1; i <= n; i++) {
    cout << dist[i] << ' ';
  }
}

int main() {
  ios::sync_with_stdio(false);
  cin.tie(0);

  cin >> n >> m >> s;
  while (m--) {
    int a, b, c;
    cin >> a >> b >> c;
    adj[a].push_back(make_pair(b, c));
  }

  dijkstra();

  return 0;
}
      `,
    };
  },
  created() {
    const id = sessionStorage.getItem('problem_id');
    this.getProblemByID(id);
  },
  methods: {
    async getProblemByID(id) {
      const { data: res } = await problemService.getProblemByID(id);
      this.problem = res.data.problem;
    },
  },
};
</script>
