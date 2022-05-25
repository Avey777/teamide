import { pattern, rule } from '@/form/base.js';

let login = {
    fields: [
        {
            label: "账号",
            name: "account",
            rules: rule.account,
        },
        {
            label: "密码",
            name: "password",
            type: "password",
            rules: rule.password,
        }
    ],
};

let register = {
    fields: [
        {
            label: "名称",
            name: "name",
            rules: [
                {
                    required: true,
                    message: `名称不能为空!`
                },
            ],
        },
        {
            label: "账号",
            name: "account",
            rules: rule.account,
        },
        {
            label: "邮箱",
            name: "email",
            type: "email",
            rules: rule.email,
        },
        {
            label: "密码",
            name: "password",
            type: "password",
            rules: rule.password,
        }
    ],
};

let app = {
    fields: [
        {
            label: "名称",
            name: "name",
            rules: [
                {
                    required: true,
                    message: `应用名称不能为空!`
                },
                {
                    pattern: /^[a-zA-Z0-9_]+$/,
                    message: `应用名称只能包含数字、字母、下划线!`
                },

            ],
        },
    ],
};


let model = {
    fields: [
        {
            label: "名称",
            name: "name",
            rules: [
                {
                    required: true,
                    message: `模型名称不能为空!`
                },
                {
                    pattern: /^[a-zA-Z0-9_/]+$/,
                    message: `模型名称只能包含数字、字母、下划线!`
                },

            ],
        },
        {
            label: "说明",
            name: "comment",
            rules: [
            ],
        },
    ],
};


let toolbox = {
    fields: [
        {
            label: "名称",
            name: "name",
            rules: [
                {
                    required: true,
                    message: `工具名称不能为空!`
                },

            ],
        },
        {
            label: "说明",
            name: "comment",
            rules: [
            ],
        },
    ],
    group: {
        fields: [
            {
                label: "名称",
                name: "name",
                rules: [
                    {
                        required: true,
                        message: `工具名称不能为空!`
                    },

                ],
            },
            {
                label: "说明",
                name: "comment",
                rules: [
                ],
            },
        ],
        option: {
            fields: [
            ],
        },
    },
};

export default {
    login,
    register,
    app,
    model,
    toolbox,
};