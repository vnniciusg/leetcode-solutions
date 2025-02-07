class Solution {
public:
    vector<int> queryResults(int limit, vector<vector<int>>& queries) {

        size_t n = queries.size();
        vector<int> result(n);
        unordered_map<int, int> color_map, ball_map;

        for (size_t i = 0; i < n; i++) {
            int ball = queries[i][0], color = queries[i][1];

            if (ball_map.find(ball) != ball_map.end()) {
                int prev_color = ball_map[ball];
                color_map[prev_color]--;

                if (color_map[prev_color] == 0) color_map.erase(prev_color);
            }

            ball_map[ball] = color;
            color_map[color]++;
            result[i] = color_map.size();
        }

        return result;
    }
};
