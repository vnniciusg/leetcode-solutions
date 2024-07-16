/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* modifiedList(vector<int>& nums, ListNode* head) {

        if (!head) return head;

        std::unordered_set<int> numSet(nums.begin(), nums.end());

        ListNode* curr = head;
        ListNode* prev = nullptr;

        while (curr) {
            if (numSet.find(curr->val) != numSet.end()) {
                if (prev) {
                    prev->next = curr->next;
                } else {
                    head = curr->next;
                }
            } else {
                prev = curr;
            }
            curr = curr->next;
        }

        return head;    
    }
};
