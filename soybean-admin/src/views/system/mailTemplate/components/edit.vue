<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { addMailTemplate, editMailTemplate } from '@/service/api/system';
import { MailTemplateOptions, translateOptions } from '@/constants/business';

defineOptions({
  name: 'MailTemplateEdit'
});

interface Props {
  /** the type of operation */
  operateType: NaiveUI.TableOperateType;
  /** the edit row data */
  rowData?: Api.System.MailTemplate | null;
}

const props = defineProps<Props>();

interface Emits {
  (e: 'submitted'): void;
}

const emit = defineEmits<Emits>();

const visible = defineModel<boolean>('visible', {
  default: false
});

const { formRef, validate, restoreValidation } = useNaiveForm();

const title = computed(() => {
  const titles: Record<NaiveUI.TableOperateType, string> = {
    add: $t('page.system.mailTemplate.addMailTemplate'),
    edit: $t('page.system.mailTemplate.editMailTemplate')
  };
  return titles[props.operateType];
});

type Model = Pick<Api.System.MailTemplate, 'name' | 'type' | 'subject' | 'contents'>;

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    name: '',
    type: 1,
    subject: '',
    contents: ''
  };
}

type RuleKey = Extract<keyof Model, 'name' | 'type' | 'subject' | 'contents'>;

const rules: Record<RuleKey, App.Global.FormRule> = {
  name: {
    required: true,
    message: $t('page.system.mailTemplate.inputName')
  },
  type: {
    required: true,
    message: $t('page.system.mailTemplate.selectType')
  },
  subject: {
    required: true,
    message: $t('page.system.mailTemplate.inputSubject')
  },
  contents: {
    required: true,
    message: $t('page.system.mailTemplate.inputContents')
  }
};

function handleInitModel() {
  Object.assign(model, createDefaultModel());

  if (props.operateType === 'edit' && props.rowData) {
    Object.assign(model, props.rowData);
  }
}

function closeDrawer() {
  visible.value = false;
}

async function handleSubmit() {
  await validate();
  const method = props.operateType === 'add' ? addMailTemplate : editMailTemplate;
  const res = await method(model);

  if (res?.error === null) {
    const k = `api.${res.message}` as App.I18n.I18nKey;
    window.$message?.success($t(k));
    closeDrawer();
    emit('submitted');
  }
}

watch(visible, () => {
  if (visible.value) {
    handleInitModel();
    restoreValidation();
  }
});
</script>

<template>
  <NModal v-model:show="visible" display-directive="show">
    <NCard :title="title" style="width: 500px">
      <NForm ref="formRef" :model="model" :rules="rules">
        <NFormItem :label="$t('dataMap.mailTemplate.name')" path="name">
          <NInput v-model:value="model.name" :placeholder="$t('page.system.mailTemplate.inputName')" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.mailTemplate.type')" path="type">
          <NSelect v-model:value="model.type" :options="translateOptions(MailTemplateOptions)" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.mailTemplate.subject')" path="subject">
          <NInput v-model:value="model.subject" :placeholder="$t('page.system.mailTemplate.inputSubject')" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.mailTemplate.contents')" path="contents">
          <NInput
            v-model:value="model.contents"
            type="textarea"
            rows="10"
            :placeholder="$t('page.system.mailTemplate.inputContents')"
          />
        </NFormItem>
      </NForm>
      <template #action>
        <NSpace :size="16" justify="end">
          <NButton @click="closeDrawer">{{ $t('common.cancel') }}</NButton>
          <NButton type="primary" @click="handleSubmit">{{ $t('common.confirm') }}</NButton>
        </NSpace>
      </template>
    </NCard>
  </NModal>
</template>

<style scoped></style>
