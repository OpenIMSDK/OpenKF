import type {
    TokenResponse,
    GetUserInfoResponse,
    UserLoginResponse,
} from '@/api/response/userModel';
import type { GetCommunityInfoResponse } from '@/api/response/communityModel';
import { defineStore } from 'pinia';
import { getMyCommunityInfo } from '@/api/index/community';
import { getMyInfo } from '@/api/index/user';

const InitUserInfo: GetUserInfoResponse = {
    uuid: '',
    email: '',
    nickname: '',
    avatar: '',
    description: '',
    is_enabled: false,
    is_admin: false,
};

const InitCommunityInfo: GetCommunityInfoResponse = {
    uuid: '',
    name: '',
    email: '',
    description: '',
    avatar: '',
};

const InitToken: TokenResponse = {
    token: '',
    expire_time_seconds: 0,
};

export const useUserStore = defineStore('user', {
    state: () => ({
        kf_token: { ...InitToken },
        im_token: { ...InitToken },
        userInfo: { ...InitUserInfo },
        communityInfo: { ...InitCommunityInfo },
    }),
    getters: {},
    actions: {
        async StoreToken(info: UserLoginResponse) {
            this.kf_token = info.kf_token;
            this.im_token = info.im_token;
        },
        async StoreInfo() {
            // fetch user info and community info
            const [userInfo, communityInfo] = await Promise.all([
                getMyInfo()
                    .then(res => {
                        console.log('getMyInfo success', res);
                        return res;
                    })
                    .catch(res => {
                        console.log('getMyInfo error', res);
                        return InitUserInfo;
                    }), // if error, set to default value
                getMyCommunityInfo()
                    .then(res => {
                        console.log('getMyCommunityInfo success', res);
                        return res;
                    })
                    .catch(res => {
                        console.log('getMyCommunityInfo error', res);
                        return InitCommunityInfo;
                    }),
            ]);
            this.userInfo = userInfo;
            this.communityInfo = communityInfo;
        },
        async logout() {
            this.kf_token = { ...InitToken };
            this.im_token = { ...InitToken };
            this.userInfo = { ...InitUserInfo };
            this.communityInfo = { ...InitCommunityInfo };
        },
    },
});
