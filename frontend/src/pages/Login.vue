<script setup lang="ts">
import { reactive } from 'vue';
import { Form, FormItem , Input, InputPassword, Button, Checkbox} from 'ant-design-vue'

interface FormState {
    username: string;
    password: string;
    remember: boolean;
}

const formState = reactive<FormState>({
    username: '',
    password: '',
    remember: true,
});
const onFinish = (values: any) => {
    console.log('Success:', values);
};

const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
};


</script>
<template>
    <div class="page-container">
        <Form size="large" :model="formState" name="basic" :label-col="{ span: 8 }" :wrapper-col="{ span: 16 }" autocomplete="off"
            @finish="onFinish" @finishFailed="onFinishFailed">
            <FormItem size="large" label="Username" name="username"
                :rules="[{ required: true, message: 'Please input your username!' }]">
                <Input size="large" v-model:value="formState.username" />
            </FormItem>

            <FormItem size="large" label="Password" name="password"
                :rules="[{ required: true, message: 'Please input your password!' }]">
                <InputPassword size="large" v-model:value="formState.password" />
            </FormItem>

            <FormItem size="large" name="remember" :wrapper-col="{ offset: 8, span: 16 }">
                <Checkbox size="large" v-model:checked="formState.remember">Remember me</Checkbox>
            </FormItem>

            <FormItem size="large" :wrapper-col="{ offset: 8, span: 16 }">
                <Button size="large" type="primary" html-type="submit">Submit</Button>
            </FormItem>
        </Form>
    </div>
</template>
<style lang="less">
.page-container {
    width: 100vw;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: white;

}
</style>
