<script lang="ts">
import TableList from "./components/TableList.svelte";
  import Tailwindcss from "./components/Tailwind.svelte";
  import Layout from "./components/Layout.svelte";
  import Modal from "./components/Modal.svelte";
  import Footer from "./components/Footer.svelte";
</script>

<main>
  <Tailwindcss />
  <Modal />
  <Layout>
	<TableList />
  </Layout>
  <Footer />
</main>
